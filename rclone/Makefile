PORTNAME=	rclone
DISTVERSIONPREFIX=	v
DISTVERSION=	1.57.0
CATEGORIES=	local net

MAINTAINER=	brad@facefault.org
COMMENT=	Sync files to and from various cloud services

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/COPYING

USES=		go:modules
USE_GITHUB=	yes
GH_ACCOUNT=	ncw

# ::v0.4.1:group_name/vendor/goftp.io/server (from gitea.com/goftp/server@v0.4.1)
#                     ~~~~~~~~~~~~~~~~~~~~~~                 aaaaa pppppp tttttt
# goftp:server:v0.4.1:vendor/goftp.io/server
# aaaaa pppppp tttttt ~~~~~~~~~~~~~~~~~~~~~~
.if exists(${.CURDIR}/Makefile.deps)
_GITEA_TUPLE!=	sed -rn 's|^[[:space:]]*\#[^/]+/(.+)[[:space:]]+\(from[[:space:]]+[^/]+/([^/]+)/([^@]+)@([^)]+).*|\2:\3:\4:\1|p' ${.CURDIR}/Makefile.deps
.else
_GITEA_TUPLE=	goftp:server:v0.4.1:vendor/goftp.io/server
.endif

PLIST_FILES=	bin/rclone

JQ_GETVER_EXPR=	.[0].name | gsub("[-a-z]"; "")

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.pre.mk>

.for account project tag subdir in ${_GITEA_TUPLE:S/:/ /g}
MASTER_SITES+=	https://gitea.com/${account}/${project}/archive/${tag}${EXTRACT_SUFX}?dummy=/:${account}_${project}
DISTFILES+=	${account}-${project}-${tag}_GITEA0${EXTRACT_SUFX}:${account}_${project}
.endfor

_OPTIONS_extract+=	750:post-extract-gitea

post-extract-gitea:
.for account project tag subdir in ${_GITEA_TUPLE:S/:/ /g}
	@${MKDIR} ${WRKSRC}/${subdir:H}
	@${RLN} ${WRKDIR}/${project} ${WRKSRC}/${subdir}
.endfor

.include <bsd.port.post.mk>
