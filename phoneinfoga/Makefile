PORTNAME=	phoneinfoga
PORTVERSION=	2.10.7
DISTVERSIONPREFIX=	v
CATEGORIES=	local

MAINTAINER=	none@freebsd.org
COMMENT=	Advanced information gathering & OSINT framework for phone numbers
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	GPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	git:devel/git \
		swag:local/swag

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	sundowndev
GH_PROJECT=	phoneinfoga

PLIST_FILES=	bin/${PORTNAME}

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/build.Version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/build.Commit=${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}'"

_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY
	${MKDIR} ${WRKSRC}/client/dist ${WRKSRC}/web/client/dist
	${TOUCH} ${WRKSRC}/client/dist/empty.txt ${WRKSRC}/web/client/dist/empty.txt
#	${RM} ${WRKSRC}/vendor/golang.org/x/tools/gopls/internal/regtest/misc/vendor_test.go
	cd ${WRKSRC} &&	${GO_CMD} generate ./...

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
