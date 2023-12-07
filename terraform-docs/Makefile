PORTNAME=	terraform-docs
PORTVERSION=	0.16.0
DISTVERSIONPREFIX=	v
PORTREVISION=	17
CATEGORIES=	local sysutils textproc

MAINTAINER=	0xdutra@gmail.com
COMMENT=	Generate documentation from Terraform modules
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	terraform-docs

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
