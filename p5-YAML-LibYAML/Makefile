PORTNAME=	YAML-LibYAML
PORTVERSION=	0.902.0
DISTVERSIONPREFIX=	v
CATEGORIES=	textproc perl5 local
MASTER_SITES=	CPAN
PKGNAMEPREFIX=	p5-

MAINTAINER=	sunpoet@FreeBSD.org
COMMENT=	Perl YAML Serialization using XS and libyaml
WWW=		https://metacpan.org/release/YAML-LibYAML

LICENSE=	ART10 GPLv1+
LICENSE_COMB=	dual
LICENSE_FILE_GPLv1+ =	${WRKSRC}/LICENSE

USES=		perl5
USE_PERL5=	configure

post-install:
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/${SITE_ARCH_REL}/auto/YAML/XS/LibYAML/LibYAML.so

.include "../Mk/${USES:Mperl*}.mk"
.include <bsd.port.mk>
