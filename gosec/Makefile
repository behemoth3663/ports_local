PORTNAME=	gosec
PORTVERSION=	2.16.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local devel security

MAINTAINER=	none@freebsd.org
COMMENT=	Golang Security Checker
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	securego

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'main.GitTag=$$(${GIT_COMMIT})' \
	-X 'main.Version=${PORTVERSION}' \
	-X 'main.GitTag=$$(${GIT_COMMIT})' \
	-X 'main.BuildDate=${%F %T:L:gmtime}'"

post-post-extract: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
