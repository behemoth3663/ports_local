PORTNAME=	plumber
PORTVERSION=	2.8.0
DISTVERSIONPREFIX=	v
#ORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Inspecting, piping, messaging, redirecting tool
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	streamdal

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/streamdal/plumber/options.VERSION=${DISTVERSIONPREFIX}${PORTVERSION}-$$(${GIT_COMMIT})'"

PLIST_FILES=	bin/${PORTNAME}

_OPTIONS_extract+=	600:pre-post-extract
pre-post-extract: .PHONY .SILENT
#	${MV} ${WRKDIR}/plz.codegen-a56dbdcddef2 ${WRKDIR}/wombat-a56dbdcddef2
#	${MV} ${WRKDIR}/go-redis-8.11.4 ${WRKDIR}/redis-8.11.4
	for i in plz.codegen:wombat go-redis:redis; do \
		IFS=':'; set -- $${i}; unset IFS; \
		${FIND} ${WRKDIR} -maxdepth 1 -type d -name "$${1}-*" | while read -r f; do \
			${MV} -v "$${f}" "${WRKDIR}/$${2}-$${f#${WRKDIR}/$${1}-}"; \
		done; \
	done

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
