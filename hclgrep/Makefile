#

PORTNAME=	hclgrep
PORTVERSION=	0.0.2
DISTVERSIONPREFIX=	v
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Syntax based grep for HCL(v2)
WWW=		https://github.com/magodo/hclgrep

LICENSE=	BSD3CLAUSE

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	magodo

GH_TAGNAME=	1b2b24c

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
