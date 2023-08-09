PORTNAME=	session-manager-plugin
PORTVERSION=	1.2.497.0
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Helps to AWS CLI with sessions to your managed instances
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aws

GO_TARGET=	./src/${PORTNAME} ./src/ssm-cli

PLIST_FILES=	${GO_TARGET:C|.*/|bin/|}

pre-deps: .PHONY .SILENT
	cd ${WRKSRC} && \
		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod init ${GO_PKGNAME} && \
		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod tidy -v
	${TAR} -C ${PATCHDIR} --exclude='go.mod' --exclude='go.sum' -cf ${.CURDIR}/${PORTNAME}.tar .
	${RM} ${PATCHDIR}/*
	${CP} -a ${WRKSRC}/go.mod ${WRKSRC}/go.sum ${PATCHDIR}

pre-deps-all: .PHONY
	${TAR} -C ${PATCHDIR} -xf ${.CURDIR}/${PORTNAME}.tar
	${RM} ${PATCH_COOKIE} ${.CURDIR}/${PORTNAME}.tar

#const Version = "1.2.0.0"
post-patch: .PHONY .SILENT
	${MV} ${WRKSRC}/src/sessionmanagerplugin-main ${WRKSRC}/src/${PORTNAME} || :
	${MV} ${WRKSRC}/src/ssmcli-main ${WRKSRC}/src/ssm-cli || :
	test -f ${WRKSRC}/vendor/github.com/${GH_ACCOUNT}/${GH_PROJECT}/src/version/version.go || exit 0; \
	${REINPLACE_CMD} -r \
		-e 's|^([[:space:]]*const[[:space:]]+Version[[:space:]]*=[[:space:]]*)".*"[[:space:]]*$$|\1"${PORTVERSION}"|' \
	${WRKSRC}/vendor/github.com/${GH_ACCOUNT}/${GH_PROJECT}/src/version/version.go; \
	${CP} -a ${PATCHDIR}/go.mod ${PATCHDIR}/go.sum ${WRKSRC}

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
