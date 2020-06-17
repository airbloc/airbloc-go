package account

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

const feePayerAPIVersion = "v1"

type FeePayer interface {
	Address(ctx context.Context) (common.Address, error)
	Transact(ctx context.Context, tx *types.Transaction) (common.Hash, error)
}

type feePayer struct {
	addr     common.Address
	client   *http.Client
	token    string
	endpoint *url.URL `json:"-"`
}

func (fpc feePayer) request(
	ctx context.Context,
	method, endpoint string,
	body io.Reader,
	expectCodes ...int,
) ([]byte, error) {
	if expectCodes == nil {
		expectCodes = []int{http.StatusOK}
	}

	endpoint = fmt.Sprintf("%s/%s%s", fpc.endpoint.String(), feePayerAPIVersion, endpoint)
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, errors.Wrap(err, "make new request")
	}
	req = req.WithContext(ctx)

	if fpc.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", fpc.token))
	}

	resp, err := fpc.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "request to fee payer")
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}
	match := false
	for _, expectCode := range expectCodes {
		if resp.StatusCode == expectCode {
			match = true
			break
		}
	}

	if !match {
		return nil, errors.Wrap(
			errors.Errorf(
				"code: %d, body: %s",
				resp.StatusCode, string(respBody)),
			"request failed",
		)
	}
	return respBody, nil
}

func (fpc feePayer) Address(ctx context.Context) (common.Address, error) {
	return fpc.addr, nil
}

func (fpc feePayer) Transact(ctx context.Context, tx *types.Transaction) (common.Hash, error) {
	rawTxData, err := tx.MarshalJSON()
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "marshal tx")
	}

	body, err := fpc.request(ctx, http.MethodPost, "/transact", bytes.NewReader(rawTxData))
	if err != nil {
		return common.Hash{}, err
	}

	var resp struct {
		TxHash common.Hash `json:"tx_hash"`
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		return common.Hash{}, errors.Wrap(err, "marshal response body")
	}
	if resp.TxHash == (common.Hash{}) {
		return common.Hash{}, errors.New("received transaction hash is empty")
	}
	return resp.TxHash, nil
}

func NewFeePayer(ctx context.Context, client *http.Client, rawurl string, token *string) (FeePayer, error) {
	if client == nil {
		client = http.DefaultClient
	}
	if token == nil {
		t := ""
		token = &t
	}

	endpoint, err := url.Parse(rawurl)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid fee payer url %s", rawurl)
	}

	fpc := &feePayer{client: client, endpoint: endpoint, token: *token}
	body, err := fpc.request(ctx, http.MethodGet, "/address", nil)
	if err != nil {
		return nil, nil
	}

	var resp struct {
		Address common.Address `json:"address"`
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, errors.Wrap(err, "marshal response body")
	}
	if resp.Address == (common.Address{}) {
		return nil, errors.New("received fee payer address is empty")
	}

	fpc.addr = resp.Address
	return fpc, nil
}
