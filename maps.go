package dumbo

import (
	"net/http"
)

type methodTyp uint

const (
	mSTUB methodTyp = 1 << iota
	mCONNECT
	mDELETE
	mGET
	mHEAD
	mOPTIONS
	mPATCH
	mPOST
	mPUT
	mTRACE
)

var mALL = mCONNECT | mDELETE | mGET | mHEAD |
	mOPTIONS | mPATCH | mPOST | mPUT | mTRACE

var methodMap = map[string]methodTyp{
	http.MethodConnect: mCONNECT,
	http.MethodDelete:  mDELETE,
	http.MethodGet:     mGET,
	http.MethodHead:    mHEAD,
	http.MethodOptions: mOPTIONS,
	http.MethodPatch:   mPATCH,
	http.MethodPost:    mPOST,
	http.MethodPut:     mPUT,
	http.MethodTrace:   mTRACE,
}

var reverseMethodMap = map[methodTyp]string{
	mCONNECT: http.MethodConnect,
	mDELETE:  http.MethodDelete,
	mGET:     http.MethodGet,
	mHEAD:    http.MethodHead,
	mOPTIONS: http.MethodOptions,
	mPATCH:   http.MethodPatch,
	mPOST:    http.MethodPost,
	mPUT:     http.MethodPut,
	mTRACE:   http.MethodTrace,
}
