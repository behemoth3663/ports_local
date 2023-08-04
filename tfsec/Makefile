PORTNAME=	tfsec
PORTVERSION=	1.28.1
DISTVERSIONPREFIX=	v
PORTREVISION=	8
CATEGORIES=	local security

MAINTAINER=	nulldutra@proton.me
COMMENT=	Security scanner for Terraform
WWW=		https://aquasecurity.github.io/tfsec

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aquasecurity

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

post-post-extract: .PHONY .SILENT

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
