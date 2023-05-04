#

PORTNAME=	checkmake
PORTVERSION=	0.2.2
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Experimental linter/analyzer for Makefiles
WWW=		https://github.com/mrtazz/checkmake

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	mrtazz

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'main.version=${PORTVERSION}' \
	-X 'main.buildTime=${%FT%TZ:L:gmtime}' \
	-X 'main.builder=FreeBSD ports' \
	-X 'main.goversion=go${GO_VERSION}'"

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
