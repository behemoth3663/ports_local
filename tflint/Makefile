# Created by: Gabriel M. Dutra <0xdutra@gmail.com>

PORTNAME=	tflint
PORTVERSION=	0.34.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	0xdutra@gmail.com
COMMENT=	Terraform Linter

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	terraform-linters

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor &&  make gomod-vendor >Makefile.deps && make makesum && mkcal
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
