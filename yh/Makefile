PORTNAME=	yh
PORTVERSION=	0.4.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	YAML syntax highlighter
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	andreazorzetto

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
