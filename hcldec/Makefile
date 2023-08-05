PORTNAME=	hcldec
PORTVERSION=	2.17.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Transforms HCL into JSON using decoding specification
WWW=		https://github.com/${GH_ACCOUNT}/hcl

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp
GH_PROJECT=	hcl

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

#const versionStr = "0.0.1-dev"
_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY
	${REINPLACE_CMD} -r \
		-e 's|^([[:space:]]*const[[:space:]]+versionStr[[:space:]]*=[[:space:]]*)".*"[[:space:]]*$$|\1"${PORTVERSION}"|' \
	${WRKSRC}/${GO_TARGET}/main.go

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
