PORTNAME=	rclone
DISTVERSIONPREFIX=	v
DISTVERSION=	1.67.0
PORTREVISION=	1
CATEGORIES=	local net

MAINTAINER=	tremere@cainites.net
COMMENT=	Sync files to and from various cloud services
WWW=		https://rclone.org/

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/COPYING

RUN_DEPENDS=	ca_root_nss>=0:security/ca_root_nss

USES=		go:modules
USE_GITHUB=	yes
GH_ACCOUNT=	rclone

# ::v0.4.1:group_name/vendor/goftp.io/server (from gitea.com/goftp/server@v0.4.1)
#                     ~~~~~~~~~~~~~~~~~~~~~~                 aaaaa pppppp tttttt
# goftp:server:v0.4.1:vendor/goftp.io/server
# aaaaa pppppp tttttt ~~~~~~~~~~~~~~~~~~~~~~
#.if exists(${.CURDIR}/Makefile.deps)
#_GITEA_TUPLE!=	sed -rn 's|^[[:space:]]*\#[^/]+/(.+)[[:space:]]+\(from[[:space:]]+[^/]+/([^/]+)/([^@]+)@([^)]+).*|\2:\3:\4:\1|p' ${.CURDIR}/Makefile.deps
#.else
_GITEA_TUPLE=	goftp:server:v2.0.1:vendor/goftp.io/server/v2
#.endif

PLIST_FILES=	bin/rclone \
		share/man/man1/rclone.1.gz

JQ_GETVER_EXPR=	.[0].name | gsub("[-a-z]"; "")

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/fs.VersionSuffix='"

OPTIONS_DEFINE=	CMOUNT

CMOUNT_DESC=	Enable cmount command (depends on fusefs-libs)

CMOUNT_LIB_DEPENDS=	libfuse.so:sysutils/fusefs-libs
CMOUNT_VARS=		GO_BUILDFLAGS+="-tags cmount"

post-install:
	${INSTALL_MAN} ${WRKSRC}/rclone.1 ${STAGEDIR}${PREFIX}/share/man/man1

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-google:go-genproto:35c7eff3a6b1:google_go_genproto/vendor/google.golang.org/genproto/googleapis/rpc \
#+google:go-genproto:35c7eff3a6b1:google_go_genproto/vendor/google.golang.org/genproto
GH_TUPLE:=	${GH_TUPLE:S|/googleapis/rpc$||1}

#-googleapis:google-cloud-go:                 v0.2.3:googleapis_google_cloud_go/vendor/cloud.google.com/go/compute/metadata
#+googleapis:google-cloud-go:compute/metadata/v0.2.3:googleapis_google_cloud_go/vendor/cloud.google.com/go
#+googleapis:google-cloud-go:compute/metadata/v0.2.3:googleapis_google_cloud_go/vendor/cloud.google.com/go/compute/metadata
GH_TUPLE:=	${GH_TUPLE:C|^(googleapis:google-cloud-go):(.*)/(compute/metadata)$|\1:\3/\2|1}

#-open-telemetry:opentelemetry-go-contrib:v0.46.1:open_telemetry_opentelemetry_go_contrib/vendor/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp \
#-open-telemetry:opentelemetry-go:v1.21.0:open_telemetry_opentelemetry_go/vendor/go.opentelemetry.io/otel
#GH_TUPLE:=	${GH_TUPLE:Nopen-telemetry\:opentelemetry-go*}
GH_TUPLE:=	${GH_TUPLE:Ncensus-instrumentation\:opencensus-go\:*:Nopen-telemetry\:opentelemetry-go*}

.include <bsd.port.pre.mk>

.for account project tag subdir in ${_GITEA_TUPLE:S/:/ /g}
MASTER_SITES+=	https://gitea.com/${account}/${project}/archive/${tag}${EXTRACT_SUFX}?dummy=/:${account}_${project}
DISTFILES+=	${account}-${project}-${tag}_GITEA0${EXTRACT_SUFX}:${account}_${project}
.endfor

_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY .SILENT
.for account project tag subdir in ${_GITEA_TUPLE:S/:/ /g}
	@${MKDIR} ${WRKSRC}/${subdir:H}
	@${RLN} ${WRKDIR}/${project} ${WRKSRC}/${subdir}
.endfor

pre-deps: .PHONY
	if [ -d '${PATCHDIR}_' ]; then \
		${RM} -r -- ${PATCHDIR}; \
	else \
		${MV} ${PATCHDIR} ${PATCHDIR}_; \
	fi

pre-deps-all: .PHONY

post-deps-all: .PHONY
	test ! -d '${PATCHDIR}_' || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include <bsd.port.post.mk>
