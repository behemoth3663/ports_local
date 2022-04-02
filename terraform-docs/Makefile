#

PORTNAME=	terraform-docs
PORTVERSION=	0.16.0
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local sysutils textproc

MAINTAINER=	0xdutra@gmail.com
COMMENT=	Generate documentation from Terraform modules

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	terraform-docs

PLIST_FILES=	bin/${PORTNAME}

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
