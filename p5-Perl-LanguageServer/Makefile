PORTNAME=	Perl-LanguageServer
PORTVERSION=	2.6.2
CATEGORIES=	devel perl5 local
MASTER_SITES=	CPAN
PKGNAMEPREFIX=	p5-

MAINTAINER=	none@freebsd.org
COMMENT=	Language Server and Debug Protocol Adapter for Perl
WWW=		https://metacpan.org/pod/Perl::LanguageServer

LICENSE=	ART20
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	\
	p5-AnyEvent>=0:devel/p5-AnyEvent \
	p5-AnyEvent-AIO>=0:devel/p5-AnyEvent-AIO \
	p5-Class-Refresh>=0:local/p5-Class-Refresh \
	p5-Compiler-Lexer>=0.23:devel/p5-Compiler-Lexer \
	p5-Coro>=0:devel/p5-Coro \
	p5-Data-Dump>=0:devel/p5-Data-Dump \
	p5-Encode-Locale>=0:converters/p5-Encode-Locale \
	p5-Hash-SafeKeys>=0:local/p5-Hash-SafeKeys \
	p5-IO-AIO>=0:devel/p5-IO-AIO \
	p5-JSON>=0:converters/p5-JSON \
	p5-Moose>=0:devel/p5-Moose \
	p5-PadWalker>=0:devel/p5-PadWalker

USES=		perl5 dos2unix
USE_PERL5=	configure

NO_ARCH=	yes

.include "../Mk/${USES:Mperl*}.mk"
.include <bsd.port.mk>
