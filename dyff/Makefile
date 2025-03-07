PORTNAME=	dyff
DISTVERSIONPREFIX=	v
DISTVERSION=	1.10.1
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	diff tool for YAML files, and sometimes JSON
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:1.24,modules

USE_GITHUB=	yes
GH_ACCOUNT=	homeport

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/internal/cmd.version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps post-deps-all: .PHONY
pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
