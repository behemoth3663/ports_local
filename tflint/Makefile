# $FreeBSD$

PORTNAME=	tflint
PORTVERSION=	0.20.2
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Terraform linter focused on possible errors, best practices, etc.

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	terraform-linters

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make go113mod-vendor && make makesum && mkcal
.sinclude "${.CURDIR}/Makefile.deps"
#.include "${.CURDIR}/../go113mod-vendor.mk"

.include <bsd.port.mk>
