PORTNAME=	certinfo
PORTVERSION=	1.0.23
DISTVERSIONPREFIX=	v
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	Print x509 certificate info
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	pete911

GO_BUILDFLAGS+=	-ldflags="-s -w -X 'main.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

_OPTIONS_patch+=	750:pre-deps-all
pre-deps-all: .PHONY .SILENT
.for _i in c go
	${SED} 's/linux/freebsd/' \
		${WRKSRC}/vendor/golang.design/x/clipboard/clipboard_linux.${_i} > \
		${WRKSRC}/vendor/golang.design/x/clipboard/clipboard_freebsd.${_i}
.endfor

pre-deps post-deps-all: .PHONY

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
