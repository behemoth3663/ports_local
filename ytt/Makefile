#

PORTNAME=	ytt
PORTVERSION=	0.43.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	YAML tool with templates and user provided values

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE
WWW=		https://github.com/k14s/ytt

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	k14s

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

WRKSRC=		${WRKDIR}/carvel-${PORTNAME}-${PORTVERSION}

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
