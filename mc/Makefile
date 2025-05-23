PORTNAME=	mc
PORTVERSION=	4.8.33
#PORTREVISION=	1
CATEGORIES=	misc shells
MASTER_SITES=	http://ftp.midnight-commander.org/ \
		https://ftp.osuosl.org/pub/midnightcommander/

MAINTAINER=	danfe@FreeBSD.org
COMMENT=	Midnight Commander, a free Norton Commander clone
WWW=		https://midnight-commander.org/

LICENSE=	GPLv3+
LICENSE_FILE=	${WRKSRC}/COPYING

RUN_DEPENDS=	zip:archivers/zip

.if ${:!diff --version 2>/dev/null | grep --count GNU\ diffutils || :!} < 1
EXTRA_PATCHES+=	${FILESDIR}/extra-patch-src_diffviewer_ydiff.c
RUN_DEPENDS+=	gdiff:textproc/diffutils
.endif

USES=		gmake gnome libtool perl5 pkgconfig python shebangfix tar:xz
SHEBANG_FILES=	${WRKSRC}/src/vfs/extfs/helpers/s3+.in \
		${WRKSRC}/src/vfs/extfs/helpers/uc1541.in
python_OLD_CMD=	@PYTHON@
USE_GNOME=	glib20
GNU_CONFIGURE=	yes
GNU_CONFIGURE_MANPREFIX=	${PREFIX}/share
CONFIGURE_ENV=	ZIP=${LOCALBASE}/bin/zip UNZIP=${UNZIP_NATIVE_CMD}

PORTDOCS=	AUTHORS FAQ HACKING MAINTAINERS NEWS README TODO

OPTIONS_DEFINE=	ASPELL DOCS EDITOR EXTATTR ICONV NLS PCRE2 SFTP SUBSHELL X11
OPTIONS_DEFAULT=	EDITOR ICONV SFTP SLANG SUBSHELL X11
OPTIONS_SINGLE=	SCREEN
OPTIONS_SINGLE_SCREEN=	SLANG NCURSES
OPTIONS_SUB=	yes

ASPELL_DESC=	Aspell support for internal editor
EDITOR_DESC=	Build with internal text editor
EXTATTR_DESC=	Extended filesystem attributes support
PCRE2_DESC=	Use PCRE2 as low-level search engine
SCREEN_DESC=	Screen library
SFTP_DESC=	Support for SFTP (via libssh)
SUBSHELL_DESC=	Build with subshell support
X11_DESC=	X11 events support (if available)

ASPELL_BUILD_DEPENDS=	${LOCALBASE}/include/aspell.h:textproc/aspell
ASPELL_RUN_DEPENDS=	${LOCALBASE}/share/aspell/en.multi:textproc/en-aspell
ASPELL_CONFIGURE_ENABLE=aspell

EDITOR_CONFIGURE_WITH=	internal-edit

EXTATTR_LIB_DEPENDS=	libe2p.so:sysutils/e2fsprogs-core

ICONV_USES=		iconv
ICONV_CONFIGURE_ENABLE=	charset

NCURSES_CONFIGURE_ON=	--with-screen=ncurses
NCURSES_CFLAGS+=	-DNCURSES_WIDECHAR

NLS_CONFIGURE_ON=	--with-libintl-prefix=${LOCALBASE}
NLS_CONFIGURE_OFF=	--disable-nls
NLS_USES=		gettext

PCRE2_LIB_DEPENDS=	libpcre2-8.so:devel/pcre2
PCRE2_CONFIGURE_ON=	--with-search-engine=pcre2

SFTP_LIB_DEPENDS=	libssh2.so:security/libssh2
SFTP_CONFIGURE_ENABLE=	vfs-sftp

SLANG_LIB_DEPENDS=	libslang.so:devel/libslang2
SLANG_CONFIGURE_ON=	--with-screen=slang

SUBSHELL_CONFIGURE_WITH=subshell
SUBSHELL_SUB_FILES=	pkg-message

X11_BUILD_DEPENDS=	${LOCALBASE}/include/X11/Xlib.h:x11/libX11
X11_CONFIGURE_WITH=	x

GH_ACCOUNT?=	MidnightCommander
GH_PROJECT?=	mc

JQ_GETVER_EXPR=	[.[].name | select(test("^[-0-9.]+$$"; "n"))][0]

post-patch:
	@${REINPLACE_CMD} -e 's|clang\*)|*)|' ${WRKSRC}/configure
	@${REINPLACE_CMD} -e 's|x$$UNRAR_VERSION.*x|$$UNRAR_VERSION -ge |' \
		${WRKSRC}/src/vfs/extfs/helpers/urar.in
	@${REINPLACE_CMD} -e '/^AM_CPPFLAGS = /s|CFLAGS)|& $$(E2P_&|' \
		${WRKSRC}/src/vfs/local/Makefile.in
	${CP} ${FILESDIR}/tf.syntax ${WRKSRC}/misc/syntax

post-patch-EXTATTR-off:
	@${REINPLACE_CMD} -e '/found_e2p=/s|yes|disabled|; /e2p library/ \
		{ s| or|,|; s|)|&, or disabled|; }' ${WRKSRC}/configure

post-install-DOCS-on:
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
	${LN} -sf ../${PORTDOCS:[1]} ${WRKSRC}/doc
	${INSTALL_DATA} ${PORTDOCS:S|^|${WRKSRC}/doc/|} ${STAGEDIR}${DOCSDIR}

post-install-X11-on:
	${INSTALL_DATA} ${FILESDIR}/${PORTNAME}.png \
		${STAGEDIR}${PREFIX}/share/pixmaps/
	${INSTALL_DATA} ${FILESDIR}/${PORTNAME}.desktop ${STAGEDIR}${DESKTOPDIR}

.include <bsd.port.mk>
