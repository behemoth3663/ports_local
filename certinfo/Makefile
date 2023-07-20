PORTNAME=	certinfo
PORTVERSION=	1.0.21
DISTVERSIONPREFIX=	v
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	Print x509 certificate info
WWW=		https://github.com/${GH_ACCOUNT}/${GH_ACCOUNT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	pete911

GO_BUILDFLAGS+=	-ldflags="-s -w -X 'main.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY .SILENT
.for _i in c go
	${SED} 's/linux/freebsd/' \
		${WRKSRC}/vendor/golang.design/x/clipboard/clipboard_linux.${_i} > \
		${WRKSRC}/vendor/golang.design/x/clipboard/clipboard_freebsd.${_i}
.endfor

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
