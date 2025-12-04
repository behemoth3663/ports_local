PORTNAME=	protoc-gen-validate
PORTVERSION=	1.3.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Protocol Buffer Validation
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	bufbuild

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

pre-deps-all pre-deps post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
