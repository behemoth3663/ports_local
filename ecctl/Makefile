PORTNAME=	ecctl
PORTVERSION=	1.15.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Official ESS and ECE command line interface
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	elastic

PLIST_FILES=	bin/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'main.version=${PORTVERSION}' \
	-X 'main.commit=$$(${GIT_COMMIT})' \
	-X 'main.built=${%FT%TZ:L:gmtime}'"

#M2T_OFFLINE=	true
#.export	M2T_OFFLINE

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	test -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
