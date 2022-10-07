#

PORTNAME=	confluence-cli
PORTVERSION=	0.2
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local net

MAINTAINER=	none@freebsd.org
COMMENT=	Command line interface to Atlassian Confluence

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	proctorlabs

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

post-patch: .PHONY
	cd ${WRKSRC} && \
		${SETENV} ${GO_ENV} ${GO_CMD} get -u=patch && \
		${SETENV} ${GO_ENV} ${GO_CMD} mod tidy -v

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
