PORTNAME=	dyff
DISTVERSIONPREFIX=	v
DISTVERSION=	1.5.8
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	diff tool for YAML files, and sometimes JSON
WWW=		https://github.com/${GH_ACCOUNT}/${GH_ACCOUNT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	homeport

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w \
		-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/internal/cmd.version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

post-post-extract: .PHONY .SILENT

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
