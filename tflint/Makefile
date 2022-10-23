# Created by: Gabriel M. Dutra <0xdutra@gmail.com>

PORTNAME=	tflint
PORTVERSION=	0.42.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	0xdutra@gmail.com
COMMENT=	Terraform Linter
WWW=		https://github.com/terraform-linters/tflint

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	terraform-linters

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
