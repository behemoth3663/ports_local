#

PORTNAME=	session-manager-plugin
PORTVERSION=	1.2.339.0
PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Helps to AWS CLI with sessions to your managed instances
WWW=		https://github.com/aws/session-manager-plugin

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aws

GO_TARGET=	./src/sessionmanagerplugin-main \
		./src/ssmcli-main

PLIST_FILES=	bin/${PORTNAME}

# github.com/aws/SSMCLI/src/message
# vendor/github.com/aws/SSMCLI/src/message/messageparser.go:170:3: cannot use nil as type uuid.UUID in return argument
# vendor/github.com/aws/SSMCLI/src/message/messageparser.go:176:3: cannot use nil as type uuid.UUID in return argument
# vendor/github.com/aws/SSMCLI/src/message/messageparser.go:182:3: cannot use nil as type uuid.UUID in return argument
# vendor/github.com/aws/SSMCLI/src/message/messageparser.go:188:3: cannot use nil as type uuid.UUID in return argument
# vendor/github.com/aws/SSMCLI/src/message/messageparser.go:194:3: cannot use nil as type uuid.UUID in return argument
# vendor/github.com/aws/SSMCLI/src/message/messageparser.go:417:11: invalid operation: input == nil (mismatched types uuid.UUID and nil)
# vendor/github.com/aws/SSMCLI/src/message/messageparser.go:497:20: undefined: uuid.CleanHyphen
# .-include "${.CURDIR}/Makefile.deps"
# GH_TUPLE=	\
# 		aws:aws-sdk-go:v1.43.21:aws_aws_sdk_go/vendor/github.com/aws/aws-sdk-go \
# 		aws:session-manager-plugin:57d7ac0605c0:aws_session_manager_plugin/vendor/github.com/aws/SSMCLI \
# 		cihub:seelog:f561c5e57575:cihub_seelog/vendor/github.com/cihub/seelog \
# 		davecgh:go-spew:v1.1.1:davecgh_go_spew/vendor/github.com/davecgh/go-spew \
# 		eiannone:keyboard:caf4b762e807:eiannone_keyboard/vendor/github.com/eiannone/keyboard \
# 		fsnotify:fsnotify:v1.5.1:fsnotify_fsnotify/vendor/github.com/fsnotify/fsnotify \
# 		go-yaml:yaml:9f266ea9e77c:go_yaml_yaml/vendor/gopkg.in/yaml.v3 \
# 		golang:crypto:3147a52a75dd:golang_crypto/vendor/golang.org/x/crypto \
# 		golang:sync:036812b2e83c:golang_sync/vendor/golang.org/x/sync \
# 		golang:sys:2edf467146b5:golang_sys/vendor/golang.org/x/sys \
# 		golang:term:03fcf44c2211:golang_term/vendor/golang.org/x/term \
# 		gorilla:websocket:v1.5.0:gorilla_websocket/vendor/github.com/gorilla/websocket \
# 		jmespath:go-jmespath:v0.4.0:jmespath_go_jmespath/vendor/github.com/jmespath/go-jmespath \
# 		myesui:uuid:v1.0.0:myesui_uuid/vendor/github.com/myesui/uuid \
# 		pmezard:go-difflib:v1.0.0:pmezard_go_difflib/vendor/github.com/pmezard/go-difflib \
# 		stretchr:objx:v0.1.0:stretchr_objx/vendor/github.com/stretchr/objx \
# 		stretchr:testify:v1.2.2:stretchr_testify_1/vendor/gopkg.in/stretchr/testify.v1 \
# 		stretchr:testify:v1.7.1:stretchr_testify/vendor/github.com/stretchr/testify \
# 		twinj:uuid:v1.0.0:twinj_uuid \
# 		xtaci:smux:v1.5.16:xtaci_smux/vendor/github.com/xtaci/smux
# post-extract:
# 	@${MKDIR} ${WRKSRC}/vendor/github.com/twinj
# 	@${RLN} ${WRKSRC_myesui_uuid} ${WRKSRC}/vendor/github.com/twinj/uuid
# _OPTIONS_extract+=	600:pre-post-extract
# pre-post-extract: .PHONY
# 	${CP} ${.CURDIR}/files/go.mod ${.CURDIR}/files/go.sum ${WRKSRC}

pre-patch: .PHONY
	${MV} ${WRKSRC}/vendor/src/* ${WRKSRC}/vendor
	${LN} -s ${WRKSRC} ${WRKSRC}/vendor/github.com/aws/SSMCLI
	cd ${WRKSRC} && \
	${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod init ${GO_PKGNAME} && \
	${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod edit \
		-require=github.com/davecgh/go-spew@v1.1.1 && \
	${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod edit \
		-require=github.com/pmezard/go-difflib@v1.0.0 && \
	${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod edit \
		-replace github.com/${GH_ACCOUNT}/SSMCLI=github.com/${GH_ACCOUNT}/${PORTNAME}@${PORTVERSION} && \
	${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy
	${MKDIR} ${WRKSRC}/vendor/github.com/davecgh \
		${WRKSRC}/vendor/github.com/pmezard
	${LN} -s ${WRKDIR}/.gopath/pkg/mod/github.com/davecgh/go-spew@v1.1.1 \
		${WRKSRC}/vendor/github.com/davecgh/go-spew
	${LN} -s ${WRKDIR}/.gopath/pkg/mod/github.com/pmezard/go-difflib@v1.0.0 \
		${WRKSRC}/vendor/github.com/pmezard/go-difflib

post-stage: .PHONY
	${MV} ${STAGEDIR}${PREFIX}/bin/sessionmanagerplugin-main ${STAGEDIR}${PREFIX}/bin/${PORTNAME}
	${MV} ${STAGEDIR}${PREFIX}/bin/ssmcli-main ${STAGEDIR}${PREFIX}/bin/ssm-cli

.include <bsd.port.mk>
