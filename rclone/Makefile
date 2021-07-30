PORTNAME=	rclone
DISTVERSIONPREFIX=	v
DISTVERSION=	1.56.0
CATEGORIES=	net

MAINTAINER=	brad@facefault.org
COMMENT=	Sync files to and from various cloud services

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/COPYING

USES=		go:modules
USE_GITHUB=	yes
GH_ACCOUNT=	ncw

PLIST_FILES=	bin/rclone

JQ_GETVER_EXPR=	.[0].name | gsub("[-a-z]"; "")

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.pre.mk>

.for account project tag subdir in ${_GITEA_TUPLE:S/:/ /g}
MASTER_SITES+=	https://gitea.com/${account}/${project}/archive/${tag}${EXTRACT_SUFX}?dummy=/:${account}_${project}
DISTFILES+=	${account}-${project}-${tag}_GITEA0${EXTRACT_SUFX}:${account}_${project}
.endfor

.include <bsd.port.post.mk>
