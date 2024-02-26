PORTNAME=	awsets
PORTVERSION=	1.1.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Utility for export all AWS resources
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/../../LICENSE

BUILD_DEPENDS=	git:devel/git

USES=		go:modules

WRKSRC_SUBDIR=	./cmd/${PORTNAME}
#GO_TARGET=	./cmd/${PORTNAME} -> main module (github.com/trek10inc/awsets) does not contain package github.com/trek10inc/awsets/cmd/awsets

USE_GITHUB=	yes
GH_ACCOUNT=	trek10inc

PLIST_FILES=	bin/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

M2T_OFFLINE=	true
.export	M2T_OFFLINE

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X main.version=${PORTVERSION} \
	-X main.commit=$$(${GIT_COMMIT}) \
	-X main.date=${%FT%TZ:L:gmtime}"

pre-deps: .PHONY
	if [ -d '${MASTERDIR}/files_' ]; then \
		${RM} -r ${MASTERDIR}/files; \
	else \
		${MV} ${MASTERDIR}/files ${MASTERDIR}/files_; \
	fi
	${RM} -r ${MASTERDIR}/files-${PORTVERSION}
	${MKDIR} ${MASTERDIR}/files-${PORTVERSION}

	cd ${WRKSRC} && \
		${CP} -a go.mod go.mod.orig && \
		${CP} -a go.sum go.sum.orig && \
		${GO_CMD} get -u=patch && \
		${GO_CMD} mod tidy && \
		${SETENV} TZ=UTC ${DIFF} -udp go.mod.orig go.mod | \
			${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${MASTERDIR}/files-${PORTVERSION}/patch-cmd_awsets_go.mod && \
		${SETENV} TZ=UTC ${DIFF} -udp go.sum.orig go.sum | \
			${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${MASTERDIR}/files-${PORTVERSION}/patch-cmd_awsets_go.sum

	cd ${WRKSRC}/../.. && \
		${CP} -a go.mod go.mod.orig && \
		${CP} -a go.sum go.sum.orig && \
		${GO_CMD} get -u=patch && \
		${GO_CMD} mod tidy && \
		${SETENV} TZ=UTC ${DIFF} -udp go.mod.orig go.mod | \
			${SED}	-e 's|^---[[:space:]]*|&../../|' -e 's|^+++[[:space:]]*|&../../|' \
				-e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${MASTERDIR}/files-${PORTVERSION}/patch-go.mod && \
		${SETENV} TZ=UTC ${DIFF} -udp go.sum.orig go.sum | \
			${SED}	-e 's|^---[[:space:]]*|&../../|' -e 's|^+++[[:space:]]*|&../../|' \
				-e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${MASTERDIR}/files-${PORTVERSION}/patch-go.sum

pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
