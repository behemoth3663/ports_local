PORTNAME=	plumber
PORTVERSION=	2.6.1
DISTVERSIONPREFIX=	v
#ORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Inspecting, piping, messaging, redirecting tool
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	streamdal

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
		-X github.com/batchcorp/${GH_PROJECT}/options.VERSION=${DISTVERSIONPREFIX}${PORTVERSION}"

PLIST_FILES=	bin/${PORTNAME}

_OPTIONS_extract+=	600:pre-post-extract
pre-post-extract: .PHONY .SILENT
	# ${MV} ${WRKDIR}/plz.codegen-a56dbdcddef2 ${WRKDIR}/wombat-a56dbdcddef2
	# ${MV} ${WRKDIR}/go-redis-8.11.4 ${WRKDIR}/redis-8.11.4
	for i in plz.codegen:wombat go-redis:redis; do \
		IFS=':'; set -- $${i}; unset IFS; \
		${FIND} ${WRKDIR} -maxdepth 1 -type d -name "$${1}-*" | while read -r f; do \
			${MV} -v "$${f}" "${WRKDIR}/$${2}-$${f#${WRKDIR}/$${1}-}"; \
		done; \
	done

#-		uber-go:thriftrw:v1.29.2:uber_go_thriftrw/vendor/go.uber.org/thriftrw \
#+		thriftrw:thriftrw-go:v1.29.2:uber_go_thriftrw/vendor/go.uber.org/thriftrw \
#It works, but:
#FATAL: /usr/ports/local/plumber/distinfo: has no SIZE record for uber-go-thriftrw-v1.29.2_GH0.tar.gz.
#GH_TUPLE:=	${GH_TUPLE:S/uber-go:thriftrw:/thriftrw:thriftrw-go:/}

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
