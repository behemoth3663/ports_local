PORTNAME=	json2csv
PORTVERSION=	1.2.1
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Converts a stream of newline separated json data to csv format
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	jehiah

GO_PKGNAME=	github.com/${GH_ACCOUNT}/${GH_PROJECT}
PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
