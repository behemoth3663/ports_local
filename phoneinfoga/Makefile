PORTNAME=	phoneinfoga
PORTVERSION=	2.11.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local

MAINTAINER=	none@freebsd.org
COMMENT=	Advanced information gathering & OSINT framework for phone numbers
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	GPLv3

BUILD_DEPENDS=	git:devel/git \
		swag:local/swag

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	sundowndev
GH_PROJECT=	phoneinfoga

PLIST_FILES=	bin/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/build.Version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/build.Commit=$$(${GIT_COMMIT})'"

_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY
	${MKDIR} ${WRKSRC}/client/dist ${WRKSRC}/web/client/dist
	${TOUCH} ${WRKSRC}/client/dist/empty.txt ${WRKSRC}/web/client/dist/empty.txt
#	${RM} ${WRKSRC}/vendor/golang.org/x/tools/gopls/internal/regtest/misc/vendor_test.go
	cd ${WRKSRC} &&	${GO_CMD} mod vendor &&	${GO_CMD} generate ./...

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
