#

PORTNAME=	tfproviderlint
PORTVERSION=	0.28.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Terraform Provider Lint Tool

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	bflad

GO_TARGET=	./cmd/tfproviderlintx

PLIST_FILES=	bin/${PORTNAME}

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.GitCommit=${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.VersionPrerelease='"

.if "${PORTNAME}" != "${GO_TARGET:T}"
post-install:
	${MV} ${STAGEDIR}${PREFIX}/bin/${GO_TARGET:T} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}
.endif

#rm -f -- distinfo Makefile.deps && doas make clean && make makesum gomod-vendor && make gomod-vendor >Makefile.deps
#doas make clean && make makesum package && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>