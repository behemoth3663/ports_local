PORTNAME=	cwl
PORTVERSION=	4.1.3
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Best way to tail AWS CloudWatch Logs from your terminal
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	lucagrulla
GH_PROJECT=	cw

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

post-patch: .PHONY .SILENT
	cd ${WRKSRC} && \
		${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod vendor
	${CHMOD} -R u+w ${WRKDIR}/.gopath

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
