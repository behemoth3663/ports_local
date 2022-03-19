#

PORTNAME=	session-manager-plugin
PORTVERSION=	1.2.295.0
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Helps to AWS CLI to start and end sessions to your managed instances

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aws

GO_TARGET=	./src/sessionmanagerplugin-main \
		./src/ssmcli-main

PLIST_FILES=	bin/${PORTNAME}

#.-include "${.CURDIR}/Makefile.deps"

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
