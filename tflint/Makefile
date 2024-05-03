# Created by: Gabriel M. Dutra <0xdutra@gmail.com>

PORTNAME=	tflint
PORTVERSION=	0.51.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	2
CATEGORIES=	local sysutils

MAINTAINER=	nulldutra@proton.me
COMMENT=	Terraform Linter
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	terraform-linters

PLIST_FILES=	bin/${PORTNAME}

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
