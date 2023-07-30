PORTNAME=	tfproviderlint
PORTVERSION=	0.29.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Terraform Provider Lint Tool
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	bflad

GO_TARGET=	./cmd/tfproviderlintx

PLIST_FILES=	bin/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.GitCommit=$$(${GIT_COMMIT})' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.VersionPrerelease='"

.if "${PORTNAME}" != "${GO_TARGET:T}"
post-install:
	${MV} ${STAGEDIR}${PREFIX}/bin/${GO_TARGET:T} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}
.endif

post-post-extract: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
