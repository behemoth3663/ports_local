PORTNAME=	phoneinfoga
PORTVERSION=	2.10.7
DISTVERSIONPREFIX=	v
CATEGORIES=	local

MAINTAINER=	none@freebsd.org
COMMENT=	Advanced information gathering & OSINT framework for phone numbers
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	GPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	sundowndev

PLIST_FILES=	bin/${PORTNAME}

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/build.Version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/build.Commit=${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}'"

#post-patch: .PHONY
#	${REINPLACE_CMD} -r \
#		-e 's/^[[:space:]]*(var[[:space:]]+Version[[:space:]]*=).*/\1 "${DISTVERSIONPREFIX}${PORTVERSION}"/' \
#		-e 's/^[[:space:]]*(var[[:space:]]+Commit[[:space:]]*=).*/\1 "f52c5338fd9d6711e5375e53a5356890680b31c7"/' \
#	${WRKSRC}/config/version.go

pre-build: .PHONY
	cd ${WRKSRC} &&	${GO_CMD} generate ./...
#	${MKDIR} ${WRKSRC}/client/dist
#	cd ${WRKSRC}/api && \
#	${GO_CMD} run github.com/jessevdk/go-assets-builder ../client/dist -o ./assets.go -p api && \
#	${GO_CMD} run ${WRKSRC}/vendor/github.com/swaggo/swag/cmd/swag/main.go init -g ./server.go --parseDependency

post-post-extract: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
