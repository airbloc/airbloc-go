package airbloc

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

const feePayerAPIVersion = "v1"

type feePayer struct {
	Address common.Address `json:"address"`
	URL     *url.URL       `json:"-"`
}

type feePayerClient struct {
	client   *http.Client
	feePayer *feePayer
}

func (fpc feePayerClient) request(
	ctx context.Context,
	method, endpoint string,
	body io.Reader,
	expectCodes ...int,
) ([]byte, error) {
	if expectCodes == nil {
		expectCodes = []int{http.StatusOK}
	}

	endpoint = path.Join(fpc.feePayer.URL.Host, feePayerAPIVersion, endpoint)
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, errors.Wrap(err, "make new request")
	}
	req = req.WithContext(ctx)

	resp, err := fpc.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "request fee payer address")
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

func (fpc feePayerClient) Address(ctx context.Context) (common.Address, error) {
	body, err := fpc.request(ctx, http.MethodGet, "address", nil)
	if err != nil {
		return common.Address{}, nil
	}

	var resp struct {
		Address common.Address `json:"address"`
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		return common.Address{}, errors.Wrap(err, "marshal response body")
	}
	return resp.Address, nil
}

func (fpc feePayerClient) Transact(ctx context.Context, tx *types.Transaction) error {
	rawTxData, err := tx.MarshalJSON()
	if err != nil {
		return errors.Wrap(err, "marshal tx")
	}

	_, err = fpc.request(ctx, http.MethodPost, "transact", bytes.NewReader(rawTxData))
	if err != nil {
		return err
	}
	return nil
}

func (fpc feePayerClient) FeePayer() feePayer {
	return *fpc.feePayer
}

func (fpc *feePayerClient) SetFeePayer(ctx context.Context, rawurl string) error {
	// fee payer url
	feePayerUrl, err := url.Parse(rawurl)
	if err != nil {
		return errors.Wrapf(err, "invalid fee payer url %s", rawurl)
	}
	fpc.feePayer.URL = feePayerUrl

	// fee payer address
	feePayerAddress, err := fpc.Address(ctx)
	if err != nil {
		return errors.Wrap(err, "fetch fee payer address")
	}
	fpc.feePayer.Address = feePayerAddress

	return nil
}
