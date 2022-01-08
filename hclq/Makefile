#

PORTNAME=	hclq
PORTVERSION=	0.5.3
#DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Command-line processor for HashiCorp Command Language files

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
#GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	mattolenik
GH_TAGNAME=	e55c48c

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -X 'github.com/${GH_ACCOUNT}/${PORTNAME}/cmd.version=${PORTVERSION}'"

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
