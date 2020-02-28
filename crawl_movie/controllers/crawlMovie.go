package controllers

import (
	"crawl_movie/models"

	"github.com/astaxie/beego"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie() {

	sMoviehtml := `
	<!DOCTYPE html>
	<html lang="zh-cmn-Hans" class="ua-linux ua-webkit">
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<meta name="renderer" content="webkit">
		<meta name="referrer" content="always">
		<meta name="google-site-verification" content="ok0wCgT20tBBgo9_zat2iAcimtN4Ftf5ccsh092Xeyw" />
		<title>
			七月与安生 (豆瓣)
	</title>
		
		<meta name="baidu-site-verification" content="cZdR4xxR7RxmM4zE" />
		<meta http-equiv="Pragma" content="no-cache">
		<meta http-equiv="Expires" content="Sun, 6 Mar 2005 01:00:00 GMT">
		
		<link rel="apple-touch-icon" href="https://img3.doubanio.com/f/movie/d59b2715fdea4968a450ee5f6c95c7d7a2030065/pics/movie/apple-touch-icon.png">
		<link href="https://img3.doubanio.com/f/shire/3e5dfc68b0f376484c50cf08a58bbca3700911dc/css/douban.css" rel="stylesheet" type="text/css">
		<link href="https://img3.doubanio.com/f/shire/ae3f5a3e3085968370b1fc63afcecb22d3284848/css/separation/_all.css" rel="stylesheet" type="text/css">
		<link href="https://img3.doubanio.com/f/movie/8864d3756094f5272d3c93e30ee2e324665855b0/css/movie/base/init.css" rel="stylesheet">
		<script type="text/javascript">var _head_start = new Date();</script>
		<script type="text/javascript" src="https://img3.doubanio.com/f/movie/0495cb173e298c28593766009c7b0a953246c5b5/js/movie/lib/jquery.js"></script>
		<script type="text/javascript" src="https://img3.doubanio.com/f/shire/5ecaf46d6954d5a30bc7d99be86ae34031646e00/js/douban.js"></script>
		<script type="text/javascript" src="https://img3.doubanio.com/f/shire/0efdc63b77f895eaf85281fb0e44d435c6239a3f/js/separation/_all.js"></script>
		
		<meta name="keywords" content="七月与安生,七月与安生,七月与安生影评,剧情介绍,电影图片,预告片,影讯,在线购票,论坛">
		<meta name="description" content="七月与安生电影简介和剧情介绍,七月与安生影评、图片、预告片、影讯、论坛、在线购票">
		<meta name="mobile-agent" content="format=html5; url=http://m.douban.com/movie/subject/25827935/"/>
		<link rel="alternate" href="android-app://com.douban.movie/doubanmovie/subject/25827935/" />
		<link rel="stylesheet" href="https://img3.doubanio.com/dae/cdnlib/libs/LikeButton/1.0.5/style.min.css">
		<script type="text/javascript" src="https://img3.doubanio.com/f/shire/77323ae72a612bba8b65f845491513ff3329b1bb/js/do.js" data-cfg-autoload="false"></script>
		<script type="text/javascript">
		  Do.add('dialog', {path: 'https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js', type: 'js'});
		  Do.add('dialog-css', {path: 'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css', type: 'css'});
		  Do.add('handlebarsjs', {path: 'https://img3.doubanio.com/f/movie/3d4f8e4a8918718256450eb6e57ec8e1f7a2e14b/js/movie/lib/handlebars.current.js', type: 'js'});
		</script>
		
	  <script type='text/javascript'>
	  var _vwo_code = (function() {
		var account_id = 249272,
		  settings_tolerance = 0,
		  library_tolerance = 2500,
		  use_existing_jquery = false,
		  // DO NOT EDIT BELOW THIS LINE
		  f=false,d=document;return{use_existing_jquery:function(){return use_existing_jquery;},library_tolerance:function(){return library_tolerance;},finish:function(){if(!f){f=true;var a=d.getElementById('_vis_opt_path_hides');if(a)a.parentNode.removeChild(a);}},finished:function(){return f;},load:function(a){var b=d.createElement('script');b.src=a;b.type='text/javascript';b.innerText;b.onerror=function(){_vwo_code.finish();};d.getElementsByTagName('head')[0].appendChild(b);},init:function(){settings_timer=setTimeout('_vwo_code.finish()',settings_tolerance);var a=d.createElement('style'),b='body{opacity:0 !important;filter:alpha(opacity=0) !important;background:none !important;}',h=d.getElementsByTagName('head')[0];a.setAttribute('id','_vis_opt_path_hides');a.setAttribute('type','text/css');if(a.styleSheet)a.styleSheet.cssText=b;else a.appendChild(d.createTextNode(b));h.appendChild(a);this.load('//dev.visualwebsiteoptimizer.com/j.php?a='+account_id+'&u='+encodeURIComponent(d.URL)+'&r='+Math.random());return settings_timer;}};}());
	
	  +function () {
		var bindEvent = function (el, type, handler) {
			var $ = window.jQuery || window.Zepto || window.$
		   if ($ && $.fn && $.fn.on) {
			   $(el).on(type, handler)
		   } else if($ && $.fn && $.fn.bind) {
			   $(el).bind(type, handler)
		   } else if (el.addEventListener){
			 el.addEventListener(type, handler, false);
		   } else if (el.attachEvent){
			 el.attachEvent("on" + type, handler);
		   } else {
			 el["on" + type] = handler;
		   }
		 }
	
		var _origin_load = _vwo_code.load
		_vwo_code.load = function () {
		  var args = [].slice.call(arguments)
		  bindEvent(window, 'load', function () {
			_origin_load.apply(_vwo_code, args)
		  })
		}
	  }()
	
	  _vwo_settings_timer = _vwo_code.init();
	  </script>
	
	
		
	
	
	<script type="application/ld+json">
	{
	  "@context": "http://schema.org",
	  "name": "七月与安生",
	  "url": "/subject/25827935/",
	  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2378140502.webp",
	  "director": 
	  [
		{
		  "@type": "Person",
		  "url": "/celebrity/1274534/",
		  "name": "曾国祥 Derek Tsang"
		}
	  ]
	,
	  "author": 
	  [
		{
		  "@type": "Person",
		  "url": "/celebrity/1359881/",
		  "name": "林咏琛 Lam Wing Sum"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1359882/",
		  "name": "李媛 Yuan Li"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1330070/",
		  "name": "许伊萌 Yimeng Xu"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1408385/",
		  "name": "吴楠 Nan Wu"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1355461/",
		  "name": "安妮宝贝 Jie Li"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1346487/",
		  "name": "许岚枫 Laam-Fung Hui"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1353704/",
		  "name": "姜博洋 Alec Keung"
		}
	  ]
	,
	  "actor": 
	  [
		{
		  "@type": "Person",
		  "url": "/celebrity/1274224/",
		  "name": "周冬雨 Dongyu Zhou"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1275243/",
		  "name": "马思纯 Sichun Ma"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1349387/",
		  "name": "李程彬 Toby Lee"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1328349/",
		  "name": "李萍 Ping Li"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1365506/",
		  "name": "蔡纲 Gang Cai"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1316715/",
		  "name": "蒙亭宜 Tingyi Meng"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1365507/",
		  "name": "沙全泽 Quanze Sha"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1330226/",
		  "name": "姚欣言 Cindy Yao"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1365508/",
		  "name": "李昊芳 Haofang Li"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1365509/",
		  "name": "蒋亭轩 Tingxuan Jiang"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1336601/",
		  "name": "陆忠 Zhong Lu"
		}
		,
		{
		  "@type": "Person",
		  "url": "/celebrity/1428167/",
		  "name": "任可 Ted"
		}
	  ]
	,
	  "datePublished": "2016-09-14",
	  "genre": ["\u5267\u60c5", "\u7231\u60c5"],
	  "duration": "PT1H49M",
	  "description": "七月（马思纯 饰）和安生（周冬雨 饰）的第一次相识在十三岁，她们一个是特立独行飞扬跋扈的“野孩子”，一个是单纯温婉循规蹈矩的“乖乖女”，从那一年开始，七月和安生几乎形影不离，她是她的光，她是她的影子，...",
	  "@type": "Movie",
	  "aggregateRating": {
		"@type": "AggregateRating",
		"ratingCount": "470867",
		"bestRating": "10",
		"worstRating": "2",
		"ratingValue": "7.6"
	  }
	}
	</script>
	
	
		<style type="text/css">
	  
	</style>
		<style type="text/css">img { max-width: 100%; }</style>
		<script type="text/javascript"></script>
		<link rel="stylesheet" href="https://img3.doubanio.com/misc/mixed_static/178ed94474cff3b8.css">
	
		<link rel="shortcut icon" href="https://img3.doubanio.com/favicon.ico" type="image/x-icon">
	</head>
	
	<body>
	  
		<script type="text/javascript">var _body_start = new Date();</script>
	
		
		
	
	
	
		<link href="//img3.doubanio.com/dae/accounts/resources/f5f3d66/shire/bundle.css" rel="stylesheet" type="text/css">
	
	
	
	<div id="db-global-nav" class="global-nav">
	  <div class="bd">
		
	<div class="top-nav-info">
	  <a href="https://accounts.douban.com/passport/login?source=movie" class="nav-login" rel="nofollow">登录/注册</a>
	</div>
	
	
		<div class="top-nav-doubanapp">
	  <a href="https://www.douban.com/doubanapp/app?channel=top-nav" class="lnk-doubanapp">下载豆瓣客户端</a>
	  <div id="doubanapp-tip">
		<a href="https://www.douban.com/doubanapp/app?channel=qipao" class="tip-link">豆瓣 <span class="version">6.0</span> 全新发布</a>
		<a href="javascript: void 0;" class="tip-close">×</a>
	  </div>
	  <div id="top-nav-appintro" class="more-items">
		<p class="appintro-title">豆瓣</p>
		<p class="qrcode">扫码直接下载</p>
		<div class="download">
		  <a href="https://www.douban.com/doubanapp/redirect?channel=top-nav&direct_dl=1&download=iOS">iPhone</a>
		  <span>·</span>
		  <a href="https://www.douban.com/doubanapp/redirect?channel=top-nav&direct_dl=1&download=Android" class="download-android">Android</a>
		</div>
	  </div>
	</div>
	
		
	
	
	<div class="global-nav-items">
	  <ul>
		<li class="">
		  <a href="https://www.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-main&quot;,&quot;uid&quot;:&quot;0&quot;}">豆瓣</a>
		</li>
		<li class="">
		  <a href="https://book.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-book&quot;,&quot;uid&quot;:&quot;0&quot;}">读书</a>
		</li>
		<li class="on">
		  <a href="https://movie.douban.com"  data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-movie&quot;,&quot;uid&quot;:&quot;0&quot;}">电影</a>
		</li>
		<li class="">
		  <a href="https://music.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-music&quot;,&quot;uid&quot;:&quot;0&quot;}">音乐</a>
		</li>
		<li class="">
		  <a href="https://www.douban.com/location" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-location&quot;,&quot;uid&quot;:&quot;0&quot;}">同城</a>
		</li>
		<li class="">
		  <a href="https://www.douban.com/group" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-group&quot;,&quot;uid&quot;:&quot;0&quot;}">小组</a>
		</li>
		<li class="">
		  <a href="https://read.douban.com&#47;?dcs=top-nav&amp;dcm=douban" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-read&quot;,&quot;uid&quot;:&quot;0&quot;}">阅读</a>
		</li>
		<li class="">
		  <a href="https://douban.fm&#47;?from_=shire_top_nav" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-fm&quot;,&quot;uid&quot;:&quot;0&quot;}">FM</a>
		</li>
		<li class="">
		  <a href="https://time.douban.com&#47;?dt_time_source=douban-web_top_nav" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-time&quot;,&quot;uid&quot;:&quot;0&quot;}">时间</a>
		</li>
		<li class="">
		  <a href="https://market.douban.com&#47;?utm_campaign=douban_top_nav&amp;utm_source=douban&amp;utm_medium=pc_web" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-market&quot;,&quot;uid&quot;:&quot;0&quot;}">豆品</a>
		</li>
		<li>
		  <a href="#more" class="bn-more"><span>更多</span></a>
		  <div class="more-items">
			<table cellpadding="0" cellspacing="0">
			  <tbody>
				<tr>
				  <td>
					<a href="https://ypy.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-ypy&quot;,&quot;uid&quot;:&quot;0&quot;}">豆瓣摄影</a>
				  </td>
				</tr>
			  </tbody>
			</table>
		  </div>
		</li>
	  </ul>
	</div>
	
	  </div>
	</div>
	<script>
	  ;window._GLOBAL_NAV = {
		DOUBAN_URL: "https://www.douban.com",
		N_NEW_NOTIS: 0,
		N_NEW_DOUMAIL: 0
	  };
	</script>
	
	
	
		<script src="//img3.doubanio.com/dae/accounts/resources/f5f3d66/shire/bundle.js" defer="defer"></script>
	
	
	
	
		
	
	
	
		<link href="//img3.doubanio.com/dae/accounts/resources/f5f3d66/movie/bundle.css" rel="stylesheet" type="text/css">
	
	
	
	
	<div id="db-nav-movie" class="nav">
	  <div class="nav-wrap">
	  <div class="nav-primary">
		<div class="nav-logo">
		  <a href="https:&#47;&#47;movie.douban.com">豆瓣电影</a>
		</div>
		<div class="nav-search">
		  <form action="https:&#47;&#47;search.douban.com&#47;movie/subject_search" method="get">
			<fieldset>
			  <legend>搜索：</legend>
			  <label for="inp-query">
			  </label>
			  <div class="inp"><input id="inp-query" name="search_text" size="22" maxlength="60" placeholder="搜索电影、电视剧、综艺、影人" value=""></div>
			  <div class="inp-btn"><input type="submit" value="搜索"></div>
			  <input type="hidden" name="cat" value="1002" />
			</fieldset>
		  </form>
		</div>
	  </div>
	  </div>
	  <div class="nav-secondary">
		
	
	<div class="nav-items">
	  <ul>
		<li    ><a href="https://movie.douban.com/cinema/nowplaying/"
		 >影讯&购票</a>
		</li>
		<li    ><a href="https://movie.douban.com/explore"
		 >选电影</a>
		</li>
		<li    ><a href="https://movie.douban.com/tv/"
		 >电视剧</a>
		</li>
		<li    ><a href="https://movie.douban.com/chart"
		 >排行榜</a>
		</li>
		<li    ><a href="https://movie.douban.com/tag/"
		 >分类</a>
		</li>
		<li    ><a href="https://movie.douban.com/review/best/"
		 >影评</a>
		</li>
		<li    ><a href="https://movie.douban.com/annual/2019?source=navigation"
		 >2019年度榜单</a>
		</li>
		<li    ><a href="https://m.douban.com/standbyme/annual2019?source=navigation"
				target="_blank"
		 >2019书影音报告</a>
		</li>
	  </ul>
	</div>
	
		<a href="https://movie.douban.com/annual/2019?source=movie_navigation" class="movieannual"></a>
	  </div>
	</div>
	
	<script id="suggResult" type="text/x-jquery-tmpl">
	  <li data-link="{{= url}}">
				<a href="{{= url}}" onclick="moreurl(this, {from:'movie_search_sugg', query:'{{= keyword }}', subject_id:'{{= id}}', i: '{{= index}}', type: '{{= type}}'})">
				<img src="{{= img}}" width="40" />
				<p>
					<em>{{= title}}</em>
					{{if year}}
						<span>{{= year}}</span>
					{{/if}}
					{{if sub_title}}
						<br /><span>{{= sub_title}}</span>
					{{/if}}
					{{if address}}
						<br /><span>{{= address}}</span>
					{{/if}}
					{{if episode}}
						{{if episode=="unknow"}}
							<br /><span>集数未知</span>
						{{else}}
							<br /><span>共{{= episode}}集</span>
						{{/if}}
					{{/if}}
				</p>
			</a>
			</li>
	  </script>
	
	
	
	
		<script src="//img3.doubanio.com/dae/accounts/resources/f5f3d66/movie/bundle.js" defer="defer"></script>
	
	
	
	
	
		
		<div id="wrapper">
			
	
			
		<div id="content">
			
	
		<div id="dale_movie_subject_top_icon"></div>
		<h1>
			<span property="v:itemreviewed">七月与安生</span>
				<span class="year">(2016)</span>
		</h1>
	
			<div class="grid-16-8 clearfix">
				
	
				
				<div class="article">
					
		
	
		
	
	
	
	
	
			<div class="indent clearfix">
				<div class="subjectwrap clearfix">
					<div class="subject clearfix">
						
	
	
	
	<div id="mainpic" class="">
		<a class="nbgnbg" href="https://movie.douban.com/subject/25827935/photos?type=R" title="点击看更多海报">
			<img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2378140502.webp" title="点击看更多海报" alt="七月与安生" rel="v:image" />
	   </a>
	</div>
	
						
	
	
	<div id="info">
			<span ><span class='pl'>导演</span>: <span class='attrs'><a href="/celebrity/1274534/" rel="v:directedBy">曾国祥</a></span></span><br/>
			<span ><span class='pl'>编剧</span>: <span class='attrs'><a href="/celebrity/1359881/">林咏琛</a> / <a href="/celebrity/1359882/">李媛</a> / <a href="/celebrity/1330070/">许伊萌</a> / <a href="/celebrity/1408385/">吴楠</a> / <a href="/celebrity/1355461/">安妮宝贝</a></span></span><br/>
			<span class="actor"><span class='pl'>主演</span>: <span class='attrs'><a href="/celebrity/1274224/" rel="v:starring">周冬雨</a> / <a href="/celebrity/1275243/" rel="v:starring">马思纯</a> / <a href="/celebrity/1349387/" rel="v:starring">李程彬</a> / <a href="/celebrity/1328349/" rel="v:starring">李萍</a> / <a href="/celebrity/1365506/" rel="v:starring">蔡纲</a> / <a href="/celebrity/1316715/" rel="v:starring">蒙亭宜</a> / <a href="/celebrity/1365507/" rel="v:starring">沙全泽</a> / <a href="/celebrity/1330226/" rel="v:starring">姚欣言</a> / <a href="/celebrity/1365508/" rel="v:starring">李昊芳</a> / <a href="/celebrity/1365509/" rel="v:starring">蒋亭轩</a></span></span><br/>
			<span class="pl">类型:</span> <span property="v:genre">剧情</span> / <span property="v:genre">爱情</span><br/>
			
			<span class="pl">制片国家/地区:</span> 中国大陆 / 中国香港<br/>
			<span class="pl">语言:</span> 汉语普通话<br/>
			<span class="pl">上映日期:</span> <span property="v:initialReleaseDate" content="2016-09-14(中国大陆)">2016-09-14(中国大陆)</span> / <span property="v:initialReleaseDate" content="2016-10-27(中国香港)">2016-10-27(中国香港)</span><br/>
			<span class="pl">片长:</span> <span property="v:runtime" content="109">109分钟</span><br/>
			<span class="pl">又名:</span> SoulMate / Soul Mate<br/>
			<span class="pl">IMDb链接:</span> <a href="https://www.imdb.com/title/tt6054290" target="_blank" rel="nofollow">tt6054290</a><br>
	
	</div>
	
	
	
	
					</div>
						
	
	
	<div id="interest_sectl">
		<div class="rating_wrap clearbox" rel="v:rating">
			<div class="clearfix">
			  <div class="rating_logo ll">豆瓣评分</div>
			  <div class="output-btn-wrap rr" style="display:none">
				<img src="https://img3.doubanio.com/f/movie/692e86756648f29457847c5cc5e161d6f6b8aaac/pics/movie/reference.png" />
				<a class="download-output-image" href="#">引用</a>
			  </div>
			  
			  
			</div>
			
	
	
	
	<div class="rating_self clearfix" typeof="v:Rating">
		<strong class="ll rating_num" property="v:average">7.6</strong>
		<span property="v:best" content="10.0"></span>
		<div class="rating_right ">
			<div class="ll bigstar bigstar40"></div>
			<div class="rating_sum">
					<a href="collections" class="rating_people">
						<span property="v:votes">470872</span>人评价
					</a>
			</div>
		</div>
	</div>
	<div class="ratings-on-weight">
		
			<div class="item">
			
			<span class="stars5 starstop" title="力荐">
				5星
			</span>
			<div class="power" style="width:24px"></div>
			<span class="rating_per">18.5%</span>
			<br />
			</div>
			<div class="item">
			
			<span class="stars4 starstop" title="推荐">
				4星
			</span>
			<div class="power" style="width:64px"></div>
			<span class="rating_per">48.1%</span>
			<br />
			</div>
			<div class="item">
			
			<span class="stars3 starstop" title="还行">
				3星
			</span>
			<div class="power" style="width:38px"></div>
			<span class="rating_per">28.9%</span>
			<br />
			</div>
			<div class="item">
			
			<span class="stars2 starstop" title="较差">
				2星
			</span>
			<div class="power" style="width:4px"></div>
			<span class="rating_per">3.5%</span>
			<br />
			</div>
			<div class="item">
			
			<span class="stars1 starstop" title="很差">
				1星
			</span>
			<div class="power" style="width:1px"></div>
			<span class="rating_per">0.9%</span>
			<br />
			</div>
	</div>
	
		</div>
			<div class="rating_betterthan">
				好于 <a href="/typerank?type_name=爱情&type=13&interval_id=75:65&action=">72% 爱情片</a><br/>
				好于 <a href="/typerank?type_name=剧情&type=11&interval_id=60:50&action=">57% 剧情片</a><br/>
			</div>
	</div>
	
	
					
				</div>
					
	
	
	
	
	<div id="interest_sect_level" class="clearfix">
			
				<a href="https://www.douban.com/reason=collectwish&amp;ck=" rel="nofollow" class="j a_show_login colbutt ll" name="pbtn-25827935-wish">
					<span>想看</span>
				</a>
				<a href="https://www.douban.com/reason=collectcollect&amp;ck=" rel="nofollow" class="j a_show_login colbutt ll" name="pbtn-25827935-collect">
					<span>看过</span>
				</a>
			<div class="ll j a_stars">
				
		
		评价:
		<span id="rating"> <span id="stars" data-solid="https://img3.doubanio.com/f/shire/5a2327c04c0c231bced131ddf3f4467eb80c1c86/pics/rating_icons/star_onmouseover.png" data-hollow="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" data-solid-2x="https://img3.doubanio.com/f/shire/7258904022439076d57303c3b06ad195bf1dc41a/pics/rating_icons/star_onmouseover@2x.png" data-hollow-2x="https://img3.doubanio.com/f/shire/95cc2fa733221bb8edd28ad56a7145a5ad33383e/pics/rating_icons/star_hollow_hover@2x.png">
	
				<a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-25827935-1">
			<img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star1" width="16" height="16"/></a>
				<a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-25827935-2">
			<img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star2" width="16" height="16"/></a>
				<a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-25827935-3">
			<img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star3" width="16" height="16"/></a>
				<a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-25827935-4">
			<img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star4" width="16" height="16"/></a>
				<a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-25827935-5">
			<img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star5" width="16" height="16"/></a>
		</span><span id="rateword" class="pl"></span>
		<input id="n_rating" type="hidden" value=""  />
		</span>
	
			</div>
		
	
	</div>
	
	
				
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	<div class="gtleft">
		<ul class="ul_subject_menu bicelink color_gray pt6 clearfix">
			
		
			
					<li> 
		<img src="https://img3.doubanio.com/f/shire/cc03d0fcf32b7ce3af7b160a0b85e5e66b47cc42/pics/short-comment.gif" />&nbsp;
			<a onclick="moreurl(this, {from:'mv_sbj_wr_cmnt_login'})" class="j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">写短评</a>
	 </li>
						<li> 
		
		<img src="https://img3.doubanio.com/f/shire/5bbf02b7b5ec12b23e214a580b6f9e481108488c/pics/add-review.gif" />&nbsp;
			<a onclick="moreurl(this, {from:'mv_sbj_wr_rv_login'})" class="j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">写影评</a>
	 </li>
					<li> 
		
	
	
	
	 </li>
					<li> 
	   
	
	   
		
		<span class="rec" id="电影-25827935">
		<a href= "#"
			data-type="电影"
			data-url="https://movie.douban.com/subject/25827935/"
			data-desc="电影《七月与安生》 (来自豆瓣) "
			data-title="电影《七月与安生》 (来自豆瓣) "
			data-pic="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2378140502.jpeg"
			class="bn-sharing ">
			分享到
		</a> &nbsp;&nbsp;
		</span>
	
		<script>
			if (!window.DoubanShareMenuList) {
				window.DoubanShareMenuList = [];
			}
			var __cache_url = __cache_url || {};
	
			(function(u){
				if(__cache_url[u]) return;
				__cache_url[u] = true;
				window.DoubanShareIcons = 'https://img3.doubanio.com/f/shire/d15ffd71f3f10a7210448fec5a68eaec66e7f7d0/pics/ic_shares.png';
	
				var initShareButton = function() {
					$.ajax({url:u,dataType:'script',cache:true});
				};
	
				if (typeof Do == 'function' && 'ready' in Do) {
					Do(
						'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css',
						'https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js',
						'https://img3.doubanio.com/f/movie/c4ab132ff4d3d64a83854c875ea79b8b541faf12/js/movie/lib/qrcode.min.js',
						initShareButton
					);
				} else if(typeof Douban == 'object' && 'loader' in Douban) {
					Douban.loader.batch(
						'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css',
						'https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js',
						'https://img3.doubanio.com/f/movie/c4ab132ff4d3d64a83854c875ea79b8b541faf12/js/movie/lib/qrcode.min.js'
					).done(initShareButton);
				}
	
			})('https://img3.doubanio.com/f/movie/32be6727ed3ad8f6c4a417d8a086355c3e7d1d27/js/movie/lib/sharebutton.js');
		</script>
	
	
	  </li>
				
	
		</ul>
	
		<script type="text/javascript">
			$(function(){
				$(".ul_subject_menu li.rec .bn-sharing").bind("click", function(){
					$.get("/blank?sbj_page_click=bn_sharing");
				});
				$(".ul_subject_menu .create_from_menu").bind("click", function(e){
					e.preventDefault();
					var $el = $(this);
					var glRoot = document.getElementById('gallery-topics-selection');
					if (window.has_gallery_topics && glRoot) {
						// 判断是否有话题
						glRoot.style.display = 'block';
					} else {
						location.href = $el.attr('href');
					}
				});
			});
		</script>
	</div>
	
	
	
	
					
	
	
	
	
		<style type="text/css">
			
	.suggestions-list li { position: relative; left: 0; top: 0; margin-bottom: 7px; height: 35px }
	.suggestions-list li .user-thumb { display: inline-block; *display: inline; float: left; margin: 2px 5px 0 0; vertical-align: top }
	.suggestions-list li .user-thumb img { width: 25px; height: 25px }
	.suggestions-list li .user-name-info { display: inline-block; *display: inline; line-height: 1.4em }
	.suggestions-list li .user-name-info .user-profile-link { color: #333; font-weight: 800 }
	.suggestions-list li .user-name-info .user-profile-link:hover { color: #4b8dc5 }
	.suggestions-list li .user-name-info p { color: #999 }
	.suggestions-list li .user-name-info b { color: #4b8dc5; font-weight: normal; cursor: pointer }
	.suggestions-list li .user-name-info b:hover { text-decoration: underline }
	.suggestions-list li .dismiss { position: absolute }
	.suggestions-list li .dismiss { color: #aaa; margin: 0 0 0 10px; top: 0; right: 0 }
	.suggestions-list li .dismiss:hover { color: #333; text-decoration: none }
	
	
	.suggest-overlay { position: absolute; z-index: 99; width: auto; background: #fff; border: 1px solid #c5c7d2;
		-moz-border-radius : 3px;
		-webkit-border-radius : 3px;
		border-radius: 3px
	}
	.suggest-overlay .bd { min-width: 220px; line-height: 1; background: #fafafa; color: #b3b3b3; padding: 5px;
		-moz-border-radius : 3px;
		-webkit-border-radius : 3px;
		border-radius: 3px
	}
	.suggest-overlay ul { color: #999; padding: 3px 0; min-width: 214px }
	.suggest-overlay li { cursor: pointer; padding: 3px 7px }
	.suggest-overlay li b { font-weight: bold }
	.suggest-overlay li .username { color: #333 }
	.suggest-overlay img { margin-right: 5px; width: 20px; height: 20px; vertical-align: middle }
	.suggest-overlay .on { background: #e9f0f8 }
	
	.mentioned-highlighter { font: 14px/20px "Helvetica Neue",Helvetica,Arial,sans-serif; position: absolute; left: 4px; top: 4px; font-size: 14px; height: 60px; width: 98.5%; overflow: hidden; background: #fff; white-space: pre-wrap; word-wrap: break-word; color: transparent }
	.mentioned-highlighter b { font-weight: normal; background-color: #d2e1f3; color: transparent;
	  -moz-border-radius: 2px;
	  -webkit-border-radius: 2px;
	  border-radius: 2px
	}
	
			.movie-share-dialog .bn-flat input {
		font-size: 14px;
	}
	.movie-share-dialog {
		z-index: 100;
	}
	.movie-share-dialog .form-ft-inner{
		text-align: right;
	}
	.movie-share-dialog div.bd {
		padding: 0;
	}
	
	.movie-share .form-bd .input-area {
		position: relative;
		margin: 15px;
		height: 91px;
		zoom: 1;
	}
	
	.movie-share-no-media .form-bd {
		height: 140px;
	}
	
	.movie-share-dialog .share-text {
		height: 85px;
		position: absolute;
		z-index: 9;
		background: transparent;
		font: 14px/18px "Helvetica Neue",Helvetica,Arial,sans-serif;
		width: 98%;
		-webkit-border-radius: 4px 4px 4px 4px;
		border-radius: 4px 4px 4px 4px;
	}
	
	.movie-share-dialog .mentioned-highlighter {
		width: 483px;
		padding: 3px 4px 4px;
		color: white;
		position: absolute;
		top:0;
		left:0;
		z-index: 0;
	}
	
	.movie-share-dialog .mentioned-highlighter code {
		color: #D2E1F3;
		background: #D2E1F3;
		border-radius: 2px;
		padding-right: 1px;
		display: inline-block;
		font: 14px/18px "Helvetica Neue",Helvetica,Arial,sans-serif;
	}
	
	
	.movie-share .form-ft {
		background: #e9eef2;
		height: 25px;
		padding-top: 10px;
		padding-bottom: 10px;
	}
	
	.movie-share .form-ft-inner {
		height: 25px;
		padding-left: 15px;
		padding-right: 15px;
	}
	
	.movie-share-dialog .dialog-only-text {
		text-align: center;
		font-size: 14px;
		line-height: 1.5;
		padding-top: 30px;
		padding-bottom: 30px;
		color: #0c7823;
	}
	
	.movie-share-dialog .ll {
		float: left;
		display: inline;
	}
	.movie-share-dialog .share-label {
		width: auto;
		display: inline;
		float: none;
	}
	
	.movie-share-dialog .leading-label {
		_vertical-align: -2px;
	}
	.movie-share-dialog .media {
		float: left;
		margin-right: 10px;
		max-width: 100px;
		max-height: 100px;
		*width: 100px;
	}
	.movie-share-dialog .info-area{
		overflow: hidden;
		zoom: 1;
		margin: 0 15px 15px;
		height: 100px;
	}
	.movie-share-dialog .info-area strong{
		font-weight: bold;
	}
	.movie-share-dialog .info-area p{
		margin: 3px 0;
		color: #999;
	}
	
	.movie-share-dialog #sync-setting {
		_vertical-align: -5px;
		margin-left: 10px;
	}
	
	.movie-share-dialog .info-area .server-error {
		position: absolute;
		bottom: 45px;
		right: 15px;
		color: red;
	}
	
	.movie-share-dialog .avail-num-indicator {
		color: #aaa;
		font-weight: 800;
		padding-right: 3px;
	}
	
	.movie-share-dialog .bottom-setting {
		width: 432px;
	}
	.movie-share-dialog .input-checkbox {
		vertical-align: -2px;
		_vertical-align: -1px;
	}
	
	.movie-share-dialog #sync-setting img {
		_vertical-align: 2px;
	}
	
	
	
	.suggest-overlay {
		z-index: 2000;
	}
	
	.movie-bar {
		position: relative;
		margin-top: 10px;
	}
	
	.movie-bar-fav {
		position: absolute;
		top: 0;
		right: 0;
	}
	
		</style>
		<script src="https://img3.doubanio.com/f/shire/a40c5220b3f40ce737b366c0030ecf810b37bfea/js/lib/mustache.js" type="text/javascript"></script>
		<script src="https://img3.doubanio.com/f/shire/1d985568f3cc434b145983919d9954e2ca627e9c/js/lib/textarea-mention.js" type="text/javascript"></script>
		<script src="https://img3.doubanio.com/f/movie/230795bbf6a9a7700cc6f395067493d5dcc572ad/js/movie/share.js" type="text/javascript"></script>
	
	<div class="rec-sec">
	<span class="rec">
		<script id="movie-share" type="text/x-html-snippet">
			
		<form class="movie-share" action="/j/share" method="POST">
			<div class="clearfix form-bd">
				<div class="input-area">
					<textarea name="text" class="share-text" cols="72" data-mention-api="https://api.douban.com/shuo/in/complete?alt=xd&amp;callback=?"></textarea>
					<input type="hidden" name="target-id" value="25827935">
					<input type="hidden" name="target-type" value="0">
					<input type="hidden" name="title" value="七月与安生‎ (2016)">
					<input type="hidden" name="desc" value="导演 曾国祥 主演 周冬雨 / 马思纯 / 中国大陆 / 中国香港 / 7.6分(470872评价)">
					<input type="hidden" name="redir" value=""/>
					<div class="mentioned-highlighter"></div>
				</div>
	
				<div class="info-area">
						<img class="media" src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2378140502.webp" />
					<strong>七月与安生‎ (2016)</strong>
					<p>导演 曾国祥 主演 周冬雨 / 马思纯 / 中国大陆 / 中国香港 / 7.6分(470872评价)</p>
					<p class="error server-error">&nbsp;</p>
				</div>
			</div>
			<div class="form-ft">
				<div class="form-ft-inner">
					
	
	
	
					<span class="avail-num-indicator">140</span>
					<span class="bn-flat">
						<input type="submit" value="推荐" />
					</span>
				</div>
			</div>
		</form>
		
		<div id="suggest-mention-tmpl" style="display:none;">
			<ul>
				{{#users}}
				<li id="{{uid}}">
				  <img src="{{avatar}}">{{{username}}}&nbsp;<span>({{{uid}}})</span>
				</li>
				{{/users}}
			</ul>
		</div>
	
	
		</script>
	
			
			<a href="/accounts/register?reason=recommend"  class="j a_show_login lnk-sharing" 
				share-id="25827935" 
				data-mode="plain" data-name="七月与安生‎ (2016)" 
				data-type="movie" data-desc="导演 曾国祥 主演 周冬雨 / 马思纯 / 中国大陆 / 中国香港 / 7.6分(470872评价)" 
				data-href="https://movie.douban.com/subject/25827935/" data-image="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2378140502.webp" 
				data-properties="{}" 
				data-redir="" data-text="" 
				data-apikey="" data-curl="" 
				data-count="10" data-object_kind="1002" 
				data-object_id="25827935" data-target_type="rec" 
				data-target_action="0" 
				data-action_props="{&#34;subject_url&#34;:&#34;https:\/\/movie.douban.com\/subject\/25827935\/&#34;,&#34;subject_title&#34;:&#34;七月与安生‎ (2016)&#34;}">推荐</a>
	</span>
	
	
	</div>
	
	
	
	
	
	
				<script type="text/javascript">
					$(function() {
						$('.collect_btn', '#interest_sect_level').each(function() {
							Douban.init_collect_btn(this);
						});
						$('html').delegate(".indent .rec-sec .lnk-sharing", "click", function() {
							moreurl(this, {
								from : 'mv_sbj_db_share'
							});
						});
					});
				</script>
			</div>
				
	
	
		<div id="collect_form_25827935"></div>
	
	
			
	
	
	
	<div class="related-info" style="margin-bottom:-10px;">
		<a name="intro"></a>
		
			
				
				
		<h2>
			<i class="">七月与安生的剧情简介</i>
				  · · · · · ·
		</h2>
	
				<div class="indent" id="link-report">
						
							<span property="v:summary" class="">
									　　七月（马思纯 饰）和安生（周冬雨 饰）的第一次相识在十三岁，她们一个是特立独行飞扬跋扈的“野孩子”，一个是单纯温婉循规蹈矩的“乖乖女”，从那一年开始，七月和安生几乎形影不离，她是她的光，她是她的影子，直到某一天，一位名为苏家明（李程彬 饰）的少年出现在了七月的身边，七月恋爱了。
										<br />
									　　安生决定前往北京讨生活，临别之前，七月意外的发现苏家明贴身带着的玉佩，竟然出现在了安生的衣领里。安生走了，七月和苏家明的恋情持续着，他们考入了同一所大学，约定一毕业就结婚。可是，事情并没有像七月所想象的那样发展，而她和苏家明之间的关系，亦因为安生的归来而产生了新的变数。
							</span>
							<span class="pl"><a href="https://movie.douban.com/help/movie#t0-qs">&copy;豆瓣</a></span>
				</div>
	</div>
	
	
		<div id="dale_movie_subject_banner_after_intro"></div>
	
		
	
	
	
	
	
	
	
	
	<div id="celebrities" class="celebrities related-celebrities">
	
	  
		<h2>
			<i class="">七月与安生的演职员</i>
				  · · · · · ·
				<span class="pl">
				(
					<a href="/subject/25827935/celebrities">全部 37</a>
				)
				</span>
		</h2>
	
	
	  <ul class="celebrities-list from-subject __oneline">
			
		
	
		
	  
	  <li class="celebrity">
		
	
	  <a href="https://movie.douban.com/celebrity/1274534/" title="曾国祥 Derek Tsang" class="">
		  <div class="avatar" style="background-image: url(https://img3.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1519542625.12.webp)">
		</div>
	  </a>
	
		<div class="info">
		  <span class="name"><a href="https://movie.douban.com/celebrity/1274534/" title="曾国祥 Derek Tsang" class="name">曾国祥</a></span>
	
		  <span class="role" title="导演">导演</span>
	
		</div>
	  </li>
	
	
			
		
	
		
	  
	  <li class="celebrity">
		
	
	  <a href="https://movie.douban.com/celebrity/1274224/" title="周冬雨 Dongyu Zhou" class="">
		  <div class="avatar" style="background-image: url(https://img1.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1481224850.68.webp)">
		</div>
	  </a>
	
		<div class="info">
		  <span class="name"><a href="https://movie.douban.com/celebrity/1274224/" title="周冬雨 Dongyu Zhou" class="name">周冬雨</a></span>
	
		  <span class="role" title="饰 李安生">饰 李安生</span>
	
		</div>
	  </li>
	
	
			
		
	
		
	  
	  <li class="celebrity">
		
	
	  <a href="https://movie.douban.com/celebrity/1275243/" title="马思纯 Sichun Ma" class="">
		  <div class="avatar" style="background-image: url(https://img1.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1477464497.27.webp)">
		</div>
	  </a>
	
		<div class="info">
		  <span class="name"><a href="https://movie.douban.com/celebrity/1275243/" title="马思纯 Sichun Ma" class="name">马思纯</a></span>
	
		  <span class="role" title="饰 林七月">饰 林七月</span>
	
		</div>
	  </li>
	
	
			
		
	
		
	  
	  <li class="celebrity">
		
	
	  <a href="https://movie.douban.com/celebrity/1349387/" title="李程彬 Toby Lee" class="">
		  <div class="avatar" style="background-image: url(https://img1.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1431402028.37.webp)">
		</div>
	  </a>
	
		<div class="info">
		  <span class="name"><a href="https://movie.douban.com/celebrity/1349387/" title="李程彬 Toby Lee" class="name">李程彬</a></span>
	
		  <span class="role" title="饰 苏家明">饰 苏家明</span>
	
		</div>
	  </li>
	
	
			
		
	
		
	  
	  <li class="celebrity">
		
	
	  <a href="https://movie.douban.com/celebrity/1328349/" title="李萍 Ping Li" class="">
		  <div class="avatar" style="background-image: url(https://img9.doubanio.com/view/celebrity/s_ratio_celebrity/public/ptkt9Jive6AAcel_avatar_uploaded1366858027.05.webp)">
		</div>
	  </a>
	
		<div class="info">
		  <span class="name"><a href="https://movie.douban.com/celebrity/1328349/" title="李萍 Ping Li" class="name">李萍</a></span>
	
		  <span class="role" title="饰 七月妈妈">饰 七月妈妈</span>
	
		</div>
	  </li>
	
	
			
		
	
		
	  
	  <li class="celebrity">
		
	
	  <a href="https://movie.douban.com/celebrity/1365506/" title="蔡纲 Gang Cai" class="">
		  <div class="avatar" style="background-image: url(https://img9.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1481202927.14.webp)">
		</div>
	  </a>
	
		<div class="info">
		  <span class="name"><a href="https://movie.douban.com/celebrity/1365506/" title="蔡纲 Gang Cai" class="name">蔡纲</a></span>
	
		  <span class="role" title="饰 七月爸爸">饰 七月爸爸</span>
	
		</div>
	  </li>
	
	
	  </ul>
	</div>
	
	
		
	
	
	<link rel="stylesheet" href="https://img3.doubanio.com/f/verify/16c7e943aee3b1dc6d65f600fcc0f6d62db7dfb4/entry_creator/dist/author_subject/style.css">
	<div id="author_subject" class="author-wrapper">
		<div class="loading"></div>
	</div>
	<script type="text/javascript">
		var answerObj = {
		  ISALL: 'False',
		  TYPE: 'movie',
		  SUBJECT_ID: '25827935',
		  USER_ID: 'None'
		}
	</script>
	<script type="text/javascript" src="https://img3.doubanio.com/f/movie/61252f2f9b35f08b37f69d17dfe48310dd295347/js/movie/lib/react/15.4/bundle.js"></script>
	<script type="text/javascript" src="https://img3.doubanio.com/f/verify/ac140ef86262b845d2be7b859e352d8196f3f6d4/entry_creator/dist/author_subject/index.js"></script>
	
	
		
	
	
	
	
	
	
	
	
	
		
		<div id="related-pic" class="related-pic">
			
		
		
		<h2>
			<i class="">七月与安生的视频和图片</i>
				  · · · · · ·
				<span class="pl">
				(
					<a href="https://movie.douban.com/subject/25827935/trailer#trailer">预告片12</a>&nbsp;|&nbsp;<a href="https://movie.douban.com/subject/25827935/trailer#short_video">视频评论1</a>&nbsp;·&nbsp;<a href="/video/create?subject_id=25827935">添加</a>&nbsp;|&nbsp;<a href="https://movie.douban.com/subject/25827935/all_photos">图片791</a>&nbsp;·&nbsp;<a href="https://movie.douban.com/subject/25827935/mupload">添加</a>
				)
				</span>
		</h2>
	
	
			<ul class="related-pic-bd  wide_videos">
					<li class="label-trailer">
						<a class="related-pic-video" href="https://movie.douban.com/trailer/203039/#content" title="预告片" style="background-image:url(https://img1.doubanio.com/img/trailer/medium/2378146297.jpg?)">
						</a>
					</li>
					
					<li class="label-short-video">
						<a class="related-pic-video" href="https://movie.douban.com/video/100597/" title="视频评论" style="background-image:url(https://img1.doubanio.com/view/photo/photo/public/p2527317978.webp?)">
						</a>
					</li>
					<li>
						<a href="https://movie.douban.com/photos/photo/2374163695/"><img src="https://img9.doubanio.com/view/photo/sqxs/public/p2374163695.webp" alt="图片" /></a>
					</li>
					<li>
						<a href="https://movie.douban.com/photos/photo/2372679263/"><img src="https://img3.doubanio.com/view/photo/sqxs/public/p2372679263.webp" alt="图片" /></a>
					</li>
			</ul>
		</div>
	
	
	
	
		
		
	
	
	
	<style type="text/css">
	.award li { display: inline; margin-right: 5px }
	.awards { margin-bottom: 20px }
	.awards h2 { background: none; color: #000; font-size: 14px; padding-bottom: 5px; margin-bottom: 8px; border-bottom: 1px dashed #dddddd }
	.awards .year { color: #666666; margin-left: -5px }
	.mod { margin-bottom: 25px }
	.mod .hd { margin-bottom: 10px }
	.mod .hd h2 {margin:24px 0 3px 0}
	</style>
	
	
	<div class="mod">
		<div class="hd">
			
		<h2>
			<i class="">七月与安生的获奖情况</i>
				  · · · · · ·
				<span class="pl">
				(
					<a href="https://movie.douban.com/subject/25827935/awards/">全部</a>
				)
				</span>
		</h2>
	
		</div>
			
			<ul class="award">
				<li>
					<a href="https://movie.douban.com/awards/goldenhorse/53/">第53届台北金马影展</a>
				</li>
				<li>金马奖 最佳导演(提名)</li>
				<li><a href='https://movie.douban.com/celebrity/1274534/' target='_blank'>曾国祥</a></li>
			</ul>
			
			<ul class="award">
				<li>
					<a href="https://movie.douban.com/awards/hkfaa/36/">第36届香港电影金像奖</a>
				</li>
				<li>最佳电影(提名)</li>
				<li></li>
			</ul>
			
			<ul class="award">
				<li>
					<a href="https://movie.douban.com/awards/golden-rooster/31/">第31届中国电影金鸡奖</a>
				</li>
				<li>最佳故事片(提名)</li>
				<li></li>
			</ul>
	</div>
	
		
	
	
	
	
	
	
	
	
		<div id="recommendations" class="">
			
			
		<h2>
			<i class="">喜欢这部电影的人也喜欢</i>
				  · · · · · ·
		</h2>
	
			
		
		<div class="recommendations-bd">
			<dl class="">
				<dt>
					<a href="https://movie.douban.com/subject/3319755/?from=subject-page" >
						<img src="https://img1.doubanio.com/view/photo/s_ratio_poster/public/p501177648.webp" alt="怦然心动" class="" />
					</a>
				</dt>
				<dd>
					<a href="https://movie.douban.com/subject/3319755/?from=subject-page" class="" >怦然心动</a>
				</dd>
			</dl>
			<dl class="">
				<dt>
					<a href="https://movie.douban.com/subject/26366465/?from=subject-page" >
						<img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2285115802.webp" alt="我的少女时代" class="" />
					</a>
				</dt>
				<dd>
					<a href="https://movie.douban.com/subject/26366465/?from=subject-page" class="" >我的少女时代</a>
				</dd>
			</dl>
			<dl class="">
				<dt>
					<a href="https://movie.douban.com/subject/30166972/?from=subject-page" >
						<img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2572166063.webp" alt="少年的你" class="" />
					</a>
				</dt>
				<dd>
					<a href="https://movie.douban.com/subject/30166972/?from=subject-page" class="" >少年的你</a>
				</dd>
			</dl>
			<dl class="">
				<dt>
					<a href="https://movie.douban.com/subject/4920528/?from=subject-page" >
						<img src="https://img9.doubanio.com/view/photo/s_ratio_poster/public/p1062824805.webp" alt="那些年，我们一起追的女孩" class="" />
					</a>
				</dt>
				<dd>
					<a href="https://movie.douban.com/subject/4920528/?from=subject-page" class="" >那些年，我们一起追的女孩</a>
				</dd>
			</dl>
			<dl class="">
				<dt>
					<a href="https://movie.douban.com/subject/4739952/?from=subject-page" >
						<img src="https://img1.doubanio.com/view/photo/s_ratio_poster/public/p767451487.webp" alt="初恋这件小事" class="" />
					</a>
				</dt>
				<dd>
					<a href="https://movie.douban.com/subject/4739952/?from=subject-page" class="" >初恋这件小事</a>
				</dd>
			</dl>
			<dl class="">
				<dt>
					<a href="https://movie.douban.com/subject/25716096/?from=subject-page" >
						<img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2540513831.webp" alt="狗十三" class="" />
					</a>
				</dt>
				<dd>
					<a href="https://movie.douban.com/subject/25716096/?from=subject-page" class="" >狗十三</a>
				</dd>
			</dl>
			<dl class="">
				<dt>
					<a href="https://movie.douban.com/subject/6874741/?from=subject-page" >
						<img src="https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2507572275.webp" alt="无问西东" class="" />
					</a>
				</dt>
				<dd>
					<a href="https://movie.douban.com/subject/6874741/?from=subject-page" class="" >无问西东</a>
				</dd>
			</dl>
			<dl class="">
				<dt>
					<a href="https://movie.douban.com/subject/11529526/?from=subject-page" >
						<img src="https://img1.doubanio.com/view/photo/s_ratio_poster/public/p1959304567.webp" alt="中国合伙人" class="" />
					</a>
				</dt>
				<dd>
					<a href="https://movie.douban.com/subject/11529526/?from=subject-page" class="" >中国合伙人</a>
				</dd>
			</dl>
			<dl class="">
				<dt>
					<a href="https://movie.douban.com/subject/27024903/?from=subject-page" >
						<img src="https://img9.doubanio.com/view/photo/s_ratio_poster/public/p2501863104.webp" alt="天才枪手" class="" />
					</a>
				</dt>
				<dd>
					<a href="https://movie.douban.com/subject/27024903/?from=subject-page" class="" >天才枪手</a>
				</dd>
			</dl>
			<dl class="">
				<dt>
					<a href="https://movie.douban.com/subject/26862829/?from=subject-page" >
						<img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2507227732.webp" alt="芳华" class="" />
					</a>
				</dt>
				<dd>
					<a href="https://movie.douban.com/subject/26862829/?from=subject-page" class="" >芳华</a>
				</dd>
			</dl>
		</div>
	
		</div>
	
	
	
			
	
	
	<script type="text/x-handlebar-tmpl" id="comment-tmpl">
		<div class="dummy-fold">
			{{#each comments}}
			<div class="comment-item" data-cid="id">
				<div class="comment">
					<h3>
						<span class="comment-vote">
								<span class="votes">{{votes}}</span>
							<input value="{{id}}" type="hidden"/>
							<a href="javascript:;" class="j {{#if ../if_logined}}a_vote_comment{{else}}a_show_login{{/if}}">有用</a>
						</span>
						<span class="comment-info">
							<a href="{{user.path}}" class="">{{user.name}}</a>
							{{#if rating}}
							<span class="allstar{{rating}}0 rating" title="{{rating_word}}"></span>
							{{/if}}
							<span>
								{{time}}
							</span>
							<p> {{content_tmpl content}} </p>
						</span>
				</div>
			</div>
			{{/each}}
		</div>
	</script>
	
	
	
	
	
	
	
	
	
	
	
	
		
	
		<div id="comments-section">
			<div class="mod-hd">
				
				
			<a class="comment_btn j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">
				<span>我要写短评</span>
			</a>
	
				
		<h2>
			<i class="">七月与安生的短评</i>
				  · · · · · ·
				<span class="pl">
				(
					<a href="https://movie.douban.com/subject/25827935/comments?status=P">全部 124093 条</a>
				)
				</span>
		</h2>
	
				
			</div>
			<div class="mod-bd">
					
	
		<div class="tab-hd">
			<a id="hot-comments-tab" href="comments" data-id="hot" class="on">热门</a>&nbsp;/&nbsp;
				<a id="new-comments-tab" href="comments?sort=time" data-id="new">最新</a>&nbsp;/&nbsp;
			<a id="following-comments-tab" href="follows_comments" data-id="following"  class="j a_show_login">好友</a>
		</div>
	
		<div class="tab-bd">
			<div id="hot-comments" class="tab">
				
		
			
			<div class="comment-item" data-cid="1084028660">
				
		
		<div class="comment">
			<h3>
				<span class="comment-vote">
					<span class="votes">3165</span>
					<input value="1084028660" type="hidden"/>
					<a href="javascript:;" class="j a_show_login" onclick="">有用</a>
				</span>
				<span class="comment-info">
					<a href="https://www.douban.com/people/MovieL/" class="">木卫二</a>
						<span>看过</span>
						<span class="allstar30 rating" title="还行"></span>
					<span class="comment-time " title="2016-09-12 22:21:02">
						2016-09-12
					</span>
				</span>
			</h3>
			<p class="">
				
					<span class="short">从问候家明开始，变得有趣了。比较有趣意外的是把作者身份也写到电影里，反转再反转还挺好看。性格与身份角色的互换，真是一出人生最好戏。当七月与安生在一起，再怎么亲密腻味你都不好察觉这份情谊的特殊与可贵。只有七月与安生被拆散，被时空阻隔，你才会发现电影要做的是逾越中国内地青春片这一类型</span>
			</p>
		</div>
	
			</div>
			
			<div class="comment-item" data-cid="1084701088">
				
		
		<div class="comment">
			<h3>
				<span class="comment-vote">
					<span class="votes">2351</span>
					<input value="1084701088" type="hidden"/>
					<a href="javascript:;" class="j a_show_login" onclick="">有用</a>
				</span>
				<span class="comment-info">
					<a href="https://www.douban.com/people/qijiuzhiyue/" class="">桃桃林林</a>
						<span>看过</span>
						<span class="allstar30 rating" title="还行"></span>
					<span class="comment-time " title="2016-09-14 19:12:34">
						2016-09-14
					</span>
				</span>
			</h3>
			<p class="">
				
					<span class="short">道理我都懂，然而就是喜欢不来。</span>
			</p>
		</div>
	
			</div>
			
			<div class="comment-item" data-cid="1083590632">
				
		
		<div class="comment">
			<h3>
				<span class="comment-vote">
					<span class="votes">2503</span>
					<input value="1083590632" type="hidden"/>
					<a href="javascript:;" class="j a_show_login" onclick="">有用</a>
				</span>
				<span class="comment-info">
					<a href="https://www.douban.com/people/shoyan/" class="">少言</a>
						<span>看过</span>
						<span class="allstar40 rating" title="推荐"></span>
					<span class="comment-time " title="2016-09-11 20:18:29">
						2016-09-11
					</span>
				</span>
			</h3>
			<p class="">
				
					<span class="short">1、好看而不俗套，走心而不尴尬，如果算作青春爱情题材，应该是近些年最好的一部；2、不敢想象周冬雨会有这样惊艳的表演，就像不敢想象曾志伟老师的儿子会把两个妞的墨迹事拍得如此到位；3、恭喜李媛和许伊萌同学，你们一直以来对写故事这件事的态度，得到了回馈；4、推荐去看，希望好片能赚钱。</span>
			</p>
		</div>
	
			</div>
			
			<div class="comment-item" data-cid="1085370135">
				
		
		<div class="comment">
			<h3>
				<span class="comment-vote">
					<span class="votes">1654</span>
					<input value="1085370135" type="hidden"/>
					<a href="javascript:;" class="j a_show_login" onclick="">有用</a>
				</span>
				<span class="comment-info">
					<a href="https://www.douban.com/people/renjiananhuo/" class="">邓安庆</a>
						<span>看过</span>
						<span class="allstar40 rating" title="推荐"></span>
					<span class="comment-time " title="2016-09-16 09:45:41">
						2016-09-16
					</span>
				</span>
			</h3>
			<p class="">
				
					<span class="short">1：剧本扎实，结构完整，女性向，那种相爱相杀互生互灭的一体感，表达得很细致；2：周冬雨演得不错啊，马思纯跟高圆圆我有点分不清了，浴室对戏很打动人；3：女性向电影，都感觉男人真是个可有可无的物种，这种蠢物不要也罢！</span>
			</p>
		</div>
	
			</div>
			
			<div class="comment-item" data-cid="1084275677">
				
		
		<div class="comment">
			<h3>
				<span class="comment-vote">
					<span class="votes">724</span>
					<input value="1084275677" type="hidden"/>
					<a href="javascript:;" class="j a_show_login" onclick="">有用</a>
				</span>
				<span class="comment-info">
					<a href="https://www.douban.com/people/duzhouchi/" class="">m89</a>
						<span>看过</span>
						<span class="allstar30 rating" title="还行"></span>
					<span class="comment-time " title="2016-09-13 16:24:48">
						2016-09-13
					</span>
				</span>
			</h3>
			<p class="">
				
					<span class="short">彭浩翔想当王家卫，曾国祥想当彭浩翔。整体偏矫情，细节很暧昧，第二次反转毫无必要。始终没想明白三个人撕逼的动机何在。安妮宝贝阿姨还是比其他网络写手多两把刷子，懂得从青春期/性的角度切入女性关系。周冬雨本色演出，马思纯（苍老师？）有点惊艳。</span>
			</p>
		</div>
	
			</div>
	
	
	
					
					&gt; <a href="comments?sort=new_score&status=P" >
						更多短评
							124093条
					</a>
			</div>
			<div id="new-comments" class="tab">
				<div id="normal">
				</div>
				<div class="fold-hd hide">
					<a class="qa" href="/help/opinion#t2-q0" target="_blank">为什么被折叠？</a>
					<a class="btn-unfold" href="#">有一些短评被折叠了</a>
					<div class="qa-tip">
						评论被折叠，是因为发布这条评论的帐号行为异常。评论仍可以被展开阅读，对发布人的账号不造成其他影响。如果认为有问题，可以<a href="https://help.douban.com/help/ask?category=movie">联系</a>豆瓣电影。
					</div>
				</div>
				<div class="fold-bd">
				</div>
				<span id="total-num"></span>
			</div>
			<div id="following-comments" class="tab">
				
		
	
	
			<div class="comment-item">
				你关注的人还没写过短评
			</div>
	
			</div>
		</div>
	
	
				
				
			</div>
		</div>
	
	
	
			
	
	<link rel="stylesheet" href="https://img3.doubanio.com/misc/mixed_static/1e4177a528d7c149.css">
	
	<section class="topics mod">
		<header>
			<h2>
				七月与安生的话题 · · · · · ·
				<span class="pl">( <span class="gallery_topics">全部 <span id="topic-count"></span> 条</span> )</span>
			</h2>
		</header>
	
		
	
	
	
	
	<section class="subject-topics">
		<div class="topic-guide" id="topic-guide">
			<img class="ic_question" src="//img3.doubanio.com/f/ithildin/b1a3edea3d04805f899e9d77c0bfc0d158df10d5/pics/export/icon_question.png">
			<div class="tip_content">
				<div class="tip_title">什么是话题</div>
				<div class="tip_desc">
					<div>无论是一部作品、一个人，还是一件事，都往往可以衍生出许多不同的话题。将这些话题细分出来，分别进行讨论，会有更多收获。</div>
				</div>
			</div>
			<img class="ic_guide" src="//img3.doubanio.com/f/ithildin/529f46d86bc08f55cd0b1843d0492242ebbd22de/pics/export/icon_guide_arrow.png">
			<img class="ic_close" id="topic-guide-close" src="//img3.doubanio.com/f/ithildin/2eb4ad488cb0854644b23f20b6fa312404429589/pics/export/close@3x.png">
		</div>
	
		<div id="topic-items"></div>
	
		<script>
			window.subject_id = 25827935;
			window.join_label_text = '写影评参与';
	
			window.topic_display_count = 4;
			window.topic_item_display_count = 1;
			window.no_content_fun_call_name = "no_topic";
	
			window.guideNode = document.getElementById('topic-guide');
			window.guideNodeClose = document.getElementById('topic-guide-close');
		</script>
		
			<link rel="stylesheet" href="https://img3.doubanio.com/f/ithildin/f731c9ea474da58c516290b3a6b1dd1237c07c5e/css/export/subject_topics.css">
			<script src="https://img3.doubanio.com/f/ithildin/d3590fc6ac47b33c804037a1aa7eec49075428c8/js/export/moment-with-locales-only-zh.js"></script>
			<script src="https://img3.doubanio.com/f/ithildin/c600fdbe69e3ffa5a3919c81ae8c8b4140e99a3e/js/export/subject_topics.js"></script>
	
	</section>
	
		<script>
			function no_topic(){
				$('#content .topics').remove();
			}
		</script>
	</section>
		<section class="reviews mod movie-content">
			<header>
				<a href="new_review" rel="nofollow" class="create-review comment_btn "
					data-isverify="False"
					data-verify-url="https://www.douban.com/accounts/phone/verify?redir=http://movie.douban.com/subject/25827935/new_review">
					<span>我要写影评</span>
				</a>
				<h2>
					七月与安生的影评 · · · · · ·
					<span class="pl">( <a href="reviews">全部 5064 条</a> )</span>
				</h2>
			</header>
	
			
	
	<style>
	#gallery-topics-selection {
	  position: fixed;
	  width: 595px;
	  padding: 40px 40px 33px 40px;
	  background: #fff;
	  border-radius: 10px;
	  box-shadow: 0 2px 16px 0 rgba(0, 0, 0, 0.2);
	  top: 50%;
	  left: 50%;
	  -webkit-transform: translate(-50%, -50%);
	  transform: translate(-50%, -50%);
	  z-index: 9999;
	}
	#gallery-topics-selection h1 {
	  font-size: 18px;
	  color: #007722;
	  margin-bottom: 36px;
	  padding: 0;
	  line-height: 28px;
	  font-weight: normal;
	}
	#gallery-topics-selection .gl_topics {
	  border-bottom: 1px solid #dfdfdf;
	  max-height: 298px;
	  overflow-y: scroll;
	}
	#gallery-topics-selection .topic {
	  margin-bottom: 24px;
	}
	#gallery-topics-selection .topic_name {
	  font-size: 15px;
	  color: #333;
	  margin: 0;
	  line-height: inherit;
	}
	#gallery-topics-selection .topic_meta {
	  font-size: 13px;
	  color: #999;
	}
	#gallery-topics-selection .topics_skip {
	  display: block;
	  cursor: pointer;
	  font-size: 16px;
	  color: #3377AA;
	  text-align: center;
	  margin-top: 33px;
	}
	#gallery-topics-selection .topics_skip:hover {
	  background: transparent;
	}
	#gallery-topics-selection .close_selection {
	  position: absolute;
	  width: 30px;
	  height: 20px;
	  top: 46px;
	  right: 40px;
	  background: #fff;
	  color: #999;
	  text-align: right;
	}
	#gallery-topics-selection .close_selection:hover{
	  background: #fff;
	  color: #999;
	}
	</style>
	
	
	
	
				<div class="review_filter">
					<a href="javascript:;;" class="cur" data-sort="">热门</a href="javascript:;;"> /
					<a href="javascript:;;" data-sort="time">最新</a href="javascript:;;"> /
					<a href="javascript:;;" data-sort="follow">好友</a href="javascript:;;">
					
				</div>
	
	
				
	
	
	
	<div class="review-list  ">
			
		
	
			
		
		<div data-cid="8076701">
			<div class="main review-item" id="8076701">
	
				
		
		<header class="main-hd">
			<a href="https://www.douban.com/people/44724062/" class="avator">
				<img width="24" height="24" src="https://img9.doubanio.com/icon/u44724062-6.jpg">
			</a>
	
			<a href="https://www.douban.com/people/44724062/" class="name">细姨妈</a>
	
				<span class="allstar50 main-title-rating" title="力荐"></span>
	
			<span content="2016-09-08" class="main-meta">2016-09-08 12:21:14</span>
	
	
		</header>
	
	
				<div class="main-bd">
	
					<h2><a href="https://movie.douban.com/review/8076701/">《Vista看天下》副主编孟静：男朋友和闺蜜出轨，隐忍不等于示弱</a></h2>
	
					<div id="review_8076701_short" class="review-short" data-rid="8076701">
						<div class="short-content">
	
							去看电影《七月与安生》，在那之前，片方的发布会找了几个嘉宾辩论，主题是：你愿意当七月还是安生？遇到闺蜜抢男友，撕还是让？  这几个嘉宾里也有我，我选了七月，被抢男友的那个。在我看来，不存在撕或者让，如果我是七月，我尽量不让男友见到闺蜜，不让他们有互留微信和电...
	
							&nbsp;(<a href="javascript:;" id="toggle-8076701-copy" class="unfold" title="展开">展开</a>)
						</div>
					</div>
	
					<div id="review_8076701_full" class="hidden">
						<div id="review_8076701_full_content" class="full-content"></div>
					</div>
	
					<div class="action">
						<a href="javascript:;" class="action-btn up" data-rid="8076701" title="有用">
							<img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
							<span id="r-useful_count-8076701">
									6628
							</span>
						</a>
						<a href="javascript:;" class="action-btn down" data-rid="8076701" title="没用">
							<img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
							<span id="r-useless_count-8076701">
									919
							</span>
						</a>
						<a href="https://movie.douban.com/review/8076701/#comments" class="reply ">693回应</a>
	
						<a href="javascript:;;" class="fold hidden">收起</a>
					</div>
				</div>
			</div>
		</div>
	
			
		
		<div data-cid="8085143">
			<div class="main review-item" id="8085143">
	
				
		
		<header class="main-hd">
			<a href="https://www.douban.com/people/60599665/" class="avator">
				<img width="24" height="24" src="https://img3.doubanio.com/icon/u60599665-40.jpg">
			</a>
	
			<a href="https://www.douban.com/people/60599665/" class="name">肖恩恩恩恩肖</a>
	
				<span class="allstar40 main-title-rating" title="推荐"></span>
	
			<span content="2016-09-14" class="main-meta">2016-09-14 21:43:16</span>
	
	
		</header>
	
	
				<div class="main-bd">
	
					<h2><a href="https://movie.douban.com/review/8085143/">喜大普奔！双女主终于不是撕【哔】机器人了</a></h2>
	
					<div id="review_8085143_short" class="review-short" data-rid="8085143">
						<div class="short-content">
	
							鉴于今年国产电影低到离奇的影片质量，在走进电影院前，我是不惮以最坏的恶意来揣测《七月与安生》的。这部电影实在太有被黑的潜力了。  让我们划一下关键词：  安妮宝贝，小清新，国产青春片，双女主，闺蜜。  全是重灾区。  《失恋三十三天》《小时代》《致青春》三部电影全...
	
							&nbsp;(<a href="javascript:;" id="toggle-8085143-copy" class="unfold" title="展开">展开</a>)
						</div>
					</div>
	
					<div id="review_8085143_full" class="hidden">
						<div id="review_8085143_full_content" class="full-content"></div>
					</div>
	
					<div class="action">
						<a href="javascript:;" class="action-btn up" data-rid="8085143" title="有用">
							<img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
							<span id="r-useful_count-8085143">
									3090
							</span>
						</a>
						<a href="javascript:;" class="action-btn down" data-rid="8085143" title="没用">
							<img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
							<span id="r-useless_count-8085143">
									398
							</span>
						</a>
						<a href="https://movie.douban.com/review/8085143/#comments" class="reply ">215回应</a>
	
						<a href="javascript:;;" class="fold hidden">收起</a>
					</div>
				</div>
			</div>
		</div>
	
			
		
		<div data-cid="8083968">
			<div class="main review-item" id="8083968">
	
				
		
		<header class="main-hd">
			<a href="https://www.douban.com/people/63773386/" class="avator">
				<img width="24" height="24" src="https://img1.doubanio.com/icon/u63773386-7.jpg">
			</a>
	
			<a href="https://www.douban.com/people/63773386/" class="name">风中芭蕾</a>
	
				<span class="allstar50 main-title-rating" title="力荐"></span>
	
			<span content="2016-09-13" class="main-meta">2016-09-13 23:05:30</span>
	
	
		</header>
	
	
				<div class="main-bd">
	
					<h2><a href="https://movie.douban.com/review/8083968/">这不是一部传统意义上的青春片，而是一部好看的剧情片。</a></h2>
	
					<div id="review_8083968_short" class="review-short" data-rid="8083968">
						<div class="short-content">
								<p class="spoiler-tip">这篇影评可能有剧透</p>
	
							自《致青春》后国产青春电影迅速成为继国产喜剧电影、爱情电影、恐怖电影之后第四大类型电影，这里说的“大”并不指生产此类电影的成熟度和标准化过程，而是指数量之多。反观影片质量，却鲜有在水准线之上的，青春电影以内容的矫情、狗血、无病呻吟、脱离现实而饱受诟病，这些...
	
							&nbsp;(<a href="javascript:;" id="toggle-8083968-copy" class="unfold" title="展开">展开</a>)
						</div>
					</div>
	
					<div id="review_8083968_full" class="hidden">
						<div id="review_8083968_full_content" class="full-content"></div>
					</div>
	
					<div class="action">
						<a href="javascript:;" class="action-btn up" data-rid="8083968" title="有用">
							<img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
							<span id="r-useful_count-8083968">
									2110
							</span>
						</a>
						<a href="javascript:;" class="action-btn down" data-rid="8083968" title="没用">
							<img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
							<span id="r-useless_count-8083968">
									208
							</span>
						</a>
						<a href="https://movie.douban.com/review/8083968/#comments" class="reply ">189回应</a>
	
						<a href="javascript:;;" class="fold hidden">收起</a>
					</div>
				</div>
			</div>
		</div>
	
			
		
		<div data-cid="8100255">
			<div class="main review-item" id="8100255">
	
				
		
		<header class="main-hd">
			<a href="https://www.douban.com/people/yysmini/" class="avator">
				<img width="24" height="24" src="https://img9.doubanio.com/icon/u2417851-44.jpg">
			</a>
	
			<a href="https://www.douban.com/people/yysmini/" class="name">雨读</a>
	
				<span class="allstar20 main-title-rating" title="较差"></span>
	
			<span content="2016-09-25" class="main-meta">2016-09-25 00:35:38</span>
	
	
		</header>
	
	
				<div class="main-bd">
	
					<h2><a href="https://movie.douban.com/review/8100255/">玉佩恒久远，一胸一撕逼</a></h2>
	
					<div id="review_8100255_short" class="review-short" data-rid="8100255">
						<div class="short-content">
	
							我是和闺蜜爱上过一个人的过来人，虽然前后时间相隔若干年，但内心里相爱相杀的戏码我懂。 我是和闺蜜一起洗澡比胸的熟练工，所以臂弯里的小温暖，离别时的小拥抱，自觉世界有她就够了，男人算个屁的独白我懂。 我的交闺蜜属性是写长长的情书，写短短的字条，写定时发送的邮件...
	
							&nbsp;(<a href="javascript:;" id="toggle-8100255-copy" class="unfold" title="展开">展开</a>)
						</div>
					</div>
	
					<div id="review_8100255_full" class="hidden">
						<div id="review_8100255_full_content" class="full-content"></div>
					</div>
	
					<div class="action">
						<a href="javascript:;" class="action-btn up" data-rid="8100255" title="有用">
							<img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
							<span id="r-useful_count-8100255">
									751
							</span>
						</a>
						<a href="javascript:;" class="action-btn down" data-rid="8100255" title="没用">
							<img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
							<span id="r-useless_count-8100255">
									129
							</span>
						</a>
						<a href="https://movie.douban.com/review/8100255/#comments" class="reply ">106回应</a>
	
						<a href="javascript:;;" class="fold hidden">收起</a>
					</div>
				</div>
			</div>
		</div>
	
			
		
		<div data-cid="8087571">
			<div class="main review-item" id="8087571">
	
				
		
		<header class="main-hd">
			<a href="https://www.douban.com/people/lephemera/" class="avator">
				<img width="24" height="24" src="https://img9.doubanio.com/icon/u2744437-144.jpg">
			</a>
	
			<a href="https://www.douban.com/people/lephemera/" class="name">蜉蝣</a>
	
				<span class="allstar30 main-title-rating" title="还行"></span>
	
			<span content="2016-09-16" class="main-meta">2016-09-16 13:30:48</span>
	
	
		</header>
	
	
				<div class="main-bd">
	
					<h2><a href="https://movie.douban.com/review/8087571/">专访 | 这是一部细节控导演拍给少女的诚意电影</a></h2>
	
					<div id="review_8087571_short" class="review-short" data-rid="8087571">
						<div class="short-content">
	
							去看《七月与安生》的时候是稍微带着点害羞的心情的，初中时候就很喜欢这部短篇，偷偷看完安妮宝贝所有的书，总有把握不好情绪被别人视为矫情的时刻。十几年后再面对这样一部少女电影，生怕它会痛杀早被更新过的神经。  于是完全没想到，在主角说出“我恨过你，但我也只有你”...
	
							&nbsp;(<a href="javascript:;" id="toggle-8087571-copy" class="unfold" title="展开">展开</a>)
						</div>
					</div>
	
					<div id="review_8087571_full" class="hidden">
						<div id="review_8087571_full_content" class="full-content"></div>
					</div>
	
					<div class="action">
						<a href="javascript:;" class="action-btn up" data-rid="8087571" title="有用">
							<img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
							<span id="r-useful_count-8087571">
									606
							</span>
						</a>
						<a href="javascript:;" class="action-btn down" data-rid="8087571" title="没用">
							<img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
							<span id="r-useless_count-8087571">
									61
							</span>
						</a>
						<a href="https://movie.douban.com/review/8087571/#comments" class="reply ">52回应</a>
	
						<a href="javascript:;;" class="fold hidden">收起</a>
					</div>
				</div>
			</div>
		</div>
	
			
		
		<div data-cid="8092536">
			<div class="main review-item" id="8092536">
	
				
		
		<header class="main-hd">
			<a href="https://www.douban.com/people/MovieL/" class="avator">
				<img width="24" height="24" src="https://img1.doubanio.com/icon/u1128221-98.jpg">
			</a>
	
			<a href="https://www.douban.com/people/MovieL/" class="name">木卫二</a>
	
				<span class="allstar30 main-title-rating" title="还行"></span>
	
			<span content="2016-09-19" class="main-meta">2016-09-19 13:21:51</span>
	
	
		</header>
	
	
				<div class="main-bd">
	
					<h2><a href="https://movie.douban.com/review/8092536/">与你无关的比赛</a></h2>
	
					<div id="review_8092536_short" class="review-short" data-rid="8092536">
						<div class="short-content">
	
							首发于微信公众号：movie432，文末有二维码噢~ 我没看过安妮宝贝的作品。 这导致看完电影，我一度分不太清楚，周冬雨和马思纯各自扮演了哪个角色。 不过，事出有因。电影版《七月与安生》，最大看点就是两个女人完成了性格与角色身份的对换。 也有一个更好记的办法。周冬雨在《...
	
							&nbsp;(<a href="javascript:;" id="toggle-8092536-copy" class="unfold" title="展开">展开</a>)
						</div>
					</div>
	
					<div id="review_8092536_full" class="hidden">
						<div id="review_8092536_full_content" class="full-content"></div>
					</div>
	
					<div class="action">
						<a href="javascript:;" class="action-btn up" data-rid="8092536" title="有用">
							<img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
							<span id="r-useful_count-8092536">
									198
							</span>
						</a>
						<a href="javascript:;" class="action-btn down" data-rid="8092536" title="没用">
							<img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
							<span id="r-useless_count-8092536">
									48
							</span>
						</a>
						<a href="https://movie.douban.com/review/8092536/#comments" class="reply ">22回应</a>
	
						<a href="javascript:;;" class="fold hidden">收起</a>
					</div>
				</div>
			</div>
		</div>
	
			
		
		<div data-cid="8087492">
			<div class="main review-item" id="8087492">
	
				
		
		<header class="main-hd">
			<a href="https://www.douban.com/people/xiaomao58/" class="avator">
				<img width="24" height="24" src="https://img1.doubanio.com/icon/u2576473-1419.jpg">
			</a>
	
			<a href="https://www.douban.com/people/xiaomao58/" class="name">毛香菇</a>
	
				<span class="allstar40 main-title-rating" title="推荐"></span>
	
			<span content="2016-09-16" class="main-meta">2016-09-16 12:27:19</span>
	
	
		</header>
	
	
				<div class="main-bd">
	
					<h2><a href="https://movie.douban.com/review/8087492/">有些人说不出哪里好 我却只想和你洗澡</a></h2>
	
					<div id="review_8087492_short" class="review-short" data-rid="8087492">
						<div class="short-content">
								<p class="spoiler-tip">这篇影评可能有剧透</p>
	
							挺小的时候看过安妮宝贝的这本原著，但对情节几乎全忘了，就只记得《七月与安生》这个名字，隐隐觉得有些百合情节。所以当得知这部电影要开拍的时候，非常期待。主创阵容出来之后，隐隐有些担心，虽然很喜欢曾国祥，但这毕竟是他第一部导演作品，两位女演员，都算无感，但看过...
	
							&nbsp;(<a href="javascript:;" id="toggle-8087492-copy" class="unfold" title="展开">展开</a>)
						</div>
					</div>
	
					<div id="review_8087492_full" class="hidden">
						<div id="review_8087492_full_content" class="full-content"></div>
					</div>
	
					<div class="action">
						<a href="javascript:;" class="action-btn up" data-rid="8087492" title="有用">
							<img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
							<span id="r-useful_count-8087492">
									396
							</span>
						</a>
						<a href="javascript:;" class="action-btn down" data-rid="8087492" title="没用">
							<img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
							<span id="r-useless_count-8087492">
									40
							</span>
						</a>
						<a href="https://movie.douban.com/review/8087492/#comments" class="reply ">51回应</a>
	
						<a href="javascript:;;" class="fold hidden">收起</a>
					</div>
				</div>
			</div>
		</div>
	
			
		
		<div data-cid="8090401">
			<div class="main review-item" id="8090401">
	
				
		
		<header class="main-hd">
			<a href="https://www.douban.com/people/sharplhl/" class="avator">
				<img width="24" height="24" src="https://img9.doubanio.com/icon/u1049139-115.jpg">
			</a>
	
			<a href="https://www.douban.com/people/sharplhl/" class="name">卢十四</a>
	
				<span class="allstar40 main-title-rating" title="推荐"></span>
	
			<span content="2016-09-18" class="main-meta">2016-09-18 01:28:25</span>
	
	
		</header>
	
	
				<div class="main-bd">
	
					<h2><a href="https://movie.douban.com/review/8090401/">七月与安生，林冲与鲁智深</a></h2>
	
					<div id="review_8090401_short" class="review-short" data-rid="8090401">
						<div class="short-content">
	
							在《七月与安生》里，两位女主角第一次撕破脸，是因为安生给七月写的那些明信片：“这么多年，你给我的每一封明信片，最后都要问候我的男朋友。”（大意）  看到这里，我马上想起了《水浒传》里那段著名的八卦：《水浒传》第五十七回，鲁智深与林冲久别重逢，鲁智深张口就问：...
	
							&nbsp;(<a href="javascript:;" id="toggle-8090401-copy" class="unfold" title="展开">展开</a>)
						</div>
					</div>
	
					<div id="review_8090401_full" class="hidden">
						<div id="review_8090401_full_content" class="full-content"></div>
					</div>
	
					<div class="action">
						<a href="javascript:;" class="action-btn up" data-rid="8090401" title="有用">
							<img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
							<span id="r-useful_count-8090401">
									196
							</span>
						</a>
						<a href="javascript:;" class="action-btn down" data-rid="8090401" title="没用">
							<img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
							<span id="r-useless_count-8090401">
									36
							</span>
						</a>
						<a href="https://movie.douban.com/review/8090401/#comments" class="reply ">35回应</a>
	
						<a href="javascript:;;" class="fold hidden">收起</a>
					</div>
				</div>
			</div>
		</div>
	
			
		
		<div data-cid="8096152">
			<div class="main review-item" id="8096152">
	
				
		
		<header class="main-hd">
			<a href="https://www.douban.com/people/itsuki1101/" class="avator">
				<img width="24" height="24" src="https://img9.doubanio.com/icon/u2211649-14.jpg">
			</a>
	
			<a href="https://www.douban.com/people/itsuki1101/" class="name">藤井树</a>
	
				<span class="allstar50 main-title-rating" title="力荐"></span>
	
			<span content="2016-09-21" class="main-meta">2016-09-21 23:17:33</span>
	
	
		</header>
	
	
				<div class="main-bd">
	
					<h2><a href="https://movie.douban.com/review/8096152/">《七月与安生》，我的热爱与使命</a></h2>
	
					<div id="review_8096152_short" class="review-short" data-rid="8096152">
						<div class="short-content">
	
							 文 |  藤井树（电影《七月与安生》总策划）   究竟是哪一年看的这篇小说，已经有点记不清了。只记得那是我买的第一本安妮宝贝的书，也是她个人出版的第一本短篇小说集《告别薇安》。那时候我还在上大学。反正一夜间安妮宝贝就红了，很多女生爱她的文字。穿棉布裙的女生，白衬...
	
							&nbsp;(<a href="javascript:;" id="toggle-8096152-copy" class="unfold" title="展开">展开</a>)
						</div>
					</div>
	
					<div id="review_8096152_full" class="hidden">
						<div id="review_8096152_full_content" class="full-content"></div>
					</div>
	
					<div class="action">
						<a href="javascript:;" class="action-btn up" data-rid="8096152" title="有用">
							<img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
							<span id="r-useful_count-8096152">
									133
							</span>
						</a>
						<a href="javascript:;" class="action-btn down" data-rid="8096152" title="没用">
							<img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
							<span id="r-useless_count-8096152">
									11
							</span>
						</a>
						<a href="https://movie.douban.com/review/8096152/#comments" class="reply ">13回应</a>
	
						<a href="javascript:;;" class="fold hidden">收起</a>
					</div>
				</div>
			</div>
		</div>
	
			
		
		<div data-cid="8209486">
			<div class="main review-item" id="8209486">
	
				
		
		<header class="main-hd">
			<a href="https://www.douban.com/people/savageecho/" class="avator">
				<img width="24" height="24" src="https://img3.doubanio.com/icon/u4001084-571.jpg">
			</a>
	
			<a href="https://www.douban.com/people/savageecho/" class="name">走走走</a>
	
				<span class="allstar50 main-title-rating" title="力荐"></span>
	
			<span content="2016-12-04" class="main-meta">2016-12-04 12:05:36</span>
	
	
		</header>
	
	
				<div class="main-bd">
	
					<h2><a href="https://movie.douban.com/review/8209486/">包子皮、床垫与难忍的男人。</a></h2>
	
					<div id="review_8209486_short" class="review-short" data-rid="8209486">
						<div class="short-content">
								<p class="spoiler-tip">这篇影评可能有剧透</p>
	
							李安生一点也不安生。  所以她离开镇江的理由，我更相信她说的，“是我自己想走的。”  苏家明当然是其中一个原因。安生发现家明喜欢自己，但七月那么喜欢家明，所以她几乎是毫不犹豫的离开镇江。就像小时候吃包子一样。七月说我就爱吃包子馅儿，安生也想吃包子馅儿，但是她夹...
	
							&nbsp;(<a href="javascript:;" id="toggle-8209486-copy" class="unfold" title="展开">展开</a>)
						</div>
					</div>
	
					<div id="review_8209486_full" class="hidden">
						<div id="review_8209486_full_content" class="full-content"></div>
					</div>
	
					<div class="action">
						<a href="javascript:;" class="action-btn up" data-rid="8209486" title="有用">
							<img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
							<span id="r-useful_count-8209486">
									173
							</span>
						</a>
						<a href="javascript:;" class="action-btn down" data-rid="8209486" title="没用">
							<img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
							<span id="r-useless_count-8209486">
									9
							</span>
						</a>
						<a href="https://movie.douban.com/review/8209486/#comments" class="reply ">22回应</a>
	
						<a href="javascript:;;" class="fold hidden">收起</a>
					</div>
				</div>
			</div>
		</div>
	
	
	
	
		
	
		
	
		<script type="text/javascript" src="https://img3.doubanio.com/misc/mixed_static/43bc338dd82b045f.js"></script>
		<!-- COLLECTED CSS -->
	</div>
	
	
	
	
	
	
	
	
					<p class="pl">
						&gt;
						<a href="reviews">
							更多影评
								5064篇
						</a>
					</p>
		</section>
	<!-- COLLECTED JS -->
	
		<br/>
	
				<div class="section-discussion">
						
						<div class="mod-hd">
								<a class="comment_btn j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow"><span>添加新讨论</span></a>
							
		<h2>
			讨论区
			 &nbsp; &middot;&nbsp; &middot;&nbsp; &middot;&nbsp; &middot;&nbsp; &middot;&nbsp; &middot;
		</h2>
	
						</div>
						
	  <table class="olt"><tr><td><td><td><td></tr>
			
			<tr>
			  <td class="pl"><a href="https://movie.douban.com/subject/25827935/discussion/614356191/" title="为什么苦等家明回头的七月，和家明睡了一觉之后就知道并介意家明最爱的不是自己">为什么苦等家明回头的七月，和家明睡了一觉之后就...</a></td>
			  <td class="pl"><span>来自</span><a href="https://www.douban.com/people/3868938/">MOON</a></td>
			  <td class="pl"><span>9 回应</span></td>
			  <td class="pl"><span>2020-02-23</span></td>
			</tr>
			
			<tr>
			  <td class="pl"><a href="https://movie.douban.com/subject/25827935/discussion/614497384/" title="为什么说其实安生与七月就是同一个人？">为什么说其实安生与七月就是同一个人？</a></td>
			  <td class="pl"><span>来自</span><a href="https://www.douban.com/people/68718557/">苒苒其木</a></td>
			  <td class="pl"><span>8 回应</span></td>
			  <td class="pl"><span>2020-02-23</span></td>
			</tr>
			
			<tr>
			  <td class="pl"><a href="https://movie.douban.com/subject/25827935/discussion/615788140/" title="重新看了一下这部电影 哭的有点头疼">重新看了一下这部电影 哭的有点头疼</a></td>
			  <td class="pl"><span>来自</span><a href="https://www.douban.com/people/M1ka/">Sisi</a></td>
			  <td class="pl"><span>1 回应</span></td>
			  <td class="pl"><span>2020-02-23</span></td>
			</tr>
			
			<tr>
			  <td class="pl"><a href="https://movie.douban.com/subject/25827935/discussion/615011692/" title="最大的败笔是最后七月的死">最大的败笔是最后七月的死</a></td>
			  <td class="pl"><span>来自</span><a href="https://www.douban.com/people/marinesj/">marinesj</a></td>
			  <td class="pl"><span>23 回应</span></td>
			  <td class="pl"><span>2020-02-23</span></td>
			</tr>
			
			<tr>
			  <td class="pl"><a href="https://movie.douban.com/subject/25827935/discussion/615808569/" title="虽然看完了，但是还是没看懂，谁是七月谁是安生啊">虽然看完了，但是还是没看懂，谁是七月谁是安生啊</a></td>
			  <td class="pl"><span>来自</span><a href="https://www.douban.com/people/184699337/">天天好心情</a></td>
			  <td class="pl"><span>2 回应</span></td>
			  <td class="pl"><span>2020-02-23</span></td>
			</tr>
	  </table>
	
						<p class="pl" align="right">
							<a href="/subject/25827935/discussion/" rel="nofollow">
								&gt; 去这部影片的讨论区（全部234条）
							</a>
						</p>
				</div>
	
			
		
			
					
	
	
	
	
	
	<div id="askmatrix">
		<div class="mod-hd">
			<h2>
				关于《七月与安生》的问题
				· · · · · ·
				<span class="pl">
					(<a href='https://movie.douban.com/subject/25827935/questions/?from=subject'>
						全部128个
					</a>)
				</span>
			</h2>
	
	
			
		
		<a class='j a_show_login comment_btn'
			href='https://movie.douban.com/subject/25827935/questions/ask/?from=subject'>我来提问</a>
	
		</div>
	
		<div class="mod-bd">
			<ul class="">
				<li class="">
					<span class="tit">
						<a href="https://movie.douban.com/subject/25827935/questions/728096/?from=subject" class="">
								家明是爱安生多一些还是爱七月多一些？
						</a>
					</span>
					<span class="meta">
						49人回答
					</span>
				</li>
				<li class="">
					<span class="tit">
						<a href="https://movie.douban.com/subject/25827935/questions/728196/?from=subject" class="">
								家明有没有跟安生出轨呢？
						</a>
					</span>
					<span class="meta">
						27人回答
					</span>
				</li>
			</ul>
	
			<p>&gt;
				<a href='https://movie.douban.com/subject/25827935/questions/?from=subject'>
					全部128个问题
				</a>
			</p>
	
		</div>
	</div>
	
	
	
				
	
	
		<script type="text/javascript">
			$(function(){if($.browser.msie && $.browser.version == 6.0){
				var $info = $('#info'),
					maxWidth = parseInt($info.css('max-width'));
				if($info.width() > maxWidth) {
					$info.width(maxWidth);
				}
			}});
		</script>
	
	
				</div>
				<div class="aside">
					
	
	
		
	
	
	
	
	
	
	
	
			
	
	
	
	
	
	
		
	
	<script id="episode-tmpl" type="text/x-jsrender">
	<div id="tv-play-source" class="play-source">
		<div class="cross">
			<span style="color:#494949; font-size:16px">{{:cn}}</span>
			<span style="cursor:pointer">✕</span>
		</div>
		<div class="episode-list">
			{{for playlist}}
				<a href="{{:play_link}}&episode={{:ep}}" target="_blank">{{:ep}}集</a>
			{{/for}}
		 <div>
	 </div>
	</script>
	
	<div class="gray_ad">
		
		<h2>
			在哪儿看这部电影
				&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;
		</h2>
	
		
		<ul class="bs">
					<li>
							<a class="playBtn" data-cn="腾讯视频" href="https://www.douban.com/link2/?url=http%3A%2F%2Fv.qq.com%2Fx%2Fcover%2Fq8742fxhp6pj7zv.html%3Fptag%3Ddouban.movie&amp;subtype=1&amp;type=online-video" target="_blank">
								腾讯视频
							</a>
						<span class="buylink-price"><span>
							免费观看
						</span></span>
					</li>
					<li>
							<a class="playBtn" data-cn="爱奇艺视频" href="https://www.douban.com/link2/?url=http%3A%2F%2Fwww.iqiyi.com%2Fv_19rr9nifyg.html%3Fvfm%3Dm_331_dbdy&amp;subtype=9&amp;type=online-video" target="_blank">
								爱奇艺视频
							</a>
						<span class="buylink-price"><span>
							免费观看
						</span></span>
					</li>
					<li>
							<a class="playBtn" data-cn="优酷视频" href="https://www.douban.com/link2/?url=http%3A%2F%2Fv.youku.com%2Fv_show%2Fid_XMTc3NTgxOTIxNg%3D%3D.html%3Ftpa%3DdW5pb25faWQ9MzAwMDA4XzEwMDAwMl8wMl8wMQ%26refer%3Desfhz_operation.xuka.xj_00003036_000000_FNZfau_19010900&amp;subtype=3&amp;type=online-video" target="_blank">
								优酷视频
							</a>
						<span class="buylink-price"><span>
							免费观看
						</span></span>
					</li>
					<li>
							<a class="playBtn" data-cn="芒果TV" href="https://www.douban.com/link2/?url=http%3A%2F%2Fwww.mgtv.com%2Fb%2F297788%2F3660147.html%3Fcxid%3D94l62sj44&amp;subtype=7&amp;type=online-video" target="_blank">
								芒果TV
							</a>
						<span class="buylink-price"><span>
							免费观看
						</span></span>
					</li>
					<li>
							<a class="playBtn" data-cn="西瓜视频" href="https://www.douban.com/link2/?url=https%3A%2F%2Fwww.ixigua.com%2Fcinema%2Falbum%2F6569374011929657864%3F%26utm_source%3Ddouban&amp;subtype=17&amp;type=online-video&amp;link2key=69c211df49" target="_blank">
								西瓜视频
							</a>
						<span class="buylink-price"><span>
							免费观看
						</span></span>
					</li>
	
		</ul>
	</div>
	
	
		<!-- douban ad begin -->
		<div id="dale_movie_subject_top_right"></div>
		<div id="dale_movie_subject_top_middle"></div>
		<!-- douban ad end -->
	
		
	
	
	
	<style type="text/css">
		.m4 {margin-bottom:8px; padding-bottom:8px;}
		.movieOnline {background:#FFF6ED; padding:10px; margin-bottom:20px;}
		.movieOnline h2 {margin:0 0 5px;}
		.movieOnline .sitename {line-height:2em; width:160px;}
		.movieOnline td,.movieOnline td a:link,.movieOnline td a:visited{color:#666;}
		.movieOnline td a:hover {color:#fff;}
		.link-bt:link,
		.link-bt:visited,
		.link-bt:hover,
		.link-bt:active {margin:5px 0 0; padding:2px 8px; background:#a8c598; color:#fff; -moz-border-radius: 3px; -webkit-border-radius: 3px; border-radius: 3px; display:inline-block;}
	</style>
	
		<div class="gray_ad">
			
		<h2>
			本片原声正在播放
				&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;
		</h2>
	
			<a target="_blank" href="https://music.douban.com/subject/26866603/">去豆瓣音乐收听</a>
	
		</div>
	
	
		
	
	
	
	
	
	
	
		
		<div class="tags">
			
			
		<h2>
			<i class="">豆瓣成员常用的标签</i>
				  · · · · · ·
		</h2>
	
			<div class="tags-body">
					<a href="/tag/青春" class="">青春</a>
					<a href="/tag/友情" class="">友情</a>
					<a href="/tag/成长" class="">成长</a>
					<a href="/tag/文艺" class="">文艺</a>
					<a href="/tag/爱情" class="">爱情</a>
					<a href="/tag/女性" class="">女性</a>
					<a href="/tag/小说改编" class="">小说改编</a>
					<a href="/tag/2016" class="">2016</a>
			</div>
		</div>
	
	
		<div id="dale_movie_review_right_guess_you_like"></div>
		<div id="dale_movie_subject_inner_middle"></div>
		<div id="dale_movie_subject_download_middle"></div>
			
	
	
	
	
	
	
	
	
	<div id="subject-doulist">
		
		
		<h2>
			<i class="">以下豆列推荐</i>
				  · · · · · ·
				<span class="pl">
				(
					<a href="https://movie.douban.com/subject/25827935/doulists">全部</a>
				)
				</span>
		</h2>
	
	
		
		<ul>
				<li>
					<a href="https://www.douban.com/doulist/30299/" target="_blank">豆瓣电影【口碑榜】2020-02-16 更新</a>
					<span>(影志)</span>
				</li>
				<li>
					<a href="https://www.douban.com/doulist/30137/" target="_blank">女子之间的情谊</a>
					<span>(煮蘑菇汤映)</span>
				</li>
				<li>
					<a href="https://www.douban.com/doulist/12747/" target="_blank">国片源流考</a>
					<span>(满头)</span>
				</li>
				<li>
					<a href="https://www.douban.com/doulist/13712178/" target="_blank">评价人数超过十万的电影（按人数排序）</a>
					<span>(依然饭特稀)</span>
				</li>
				<li>
					<a href="https://www.douban.com/doulist/257125/" target="_blank">我爱的文艺片（3）</a>
					<span>(我爱卡姆)</span>
				</li>
		</ul>
	
	</div>
	
				
	
	
	
	
	
	
	
	
	<div id="subject-others-interests">
		
		
		<h2>
			<i class="">谁在看这部电影</i>
				  · · · · · ·
		</h2>
	
		
		<ul class="">
				
				<li class="">
					<a href="https://www.douban.com/people/203655782/" class="others-interest-avatar">
						<img src="https://img9.doubanio.com/icon/u203655782-4.jpg" class="pil" alt="爱读书的人">
					</a>
					<div class="others-interest-info">
						<a href="https://www.douban.com/people/203655782/" class="">爱读书的人</a>
						<div class="">
							1分钟前
							看过
							
						</div>
					</div>
				</li>
				
				<li class="">
					<a href="https://www.douban.com/people/188838608/" class="others-interest-avatar">
						<img src="https://img3.doubanio.com/icon/u188838608-1.jpg" class="pil" alt="清欢">
					</a>
					<div class="others-interest-info">
						<a href="https://www.douban.com/people/188838608/" class="">清欢</a>
						<div class="">
							1分钟前
							看过
							<span class="allstar50" title="力荐"></span>
						</div>
					</div>
				</li>
				
				<li class="">
					<a href="https://www.douban.com/people/212097281/" class="others-interest-avatar">
						<img src="https://img3.doubanio.com/icon/u212097281-1.jpg" class="pil" alt="bananana0">
					</a>
					<div class="others-interest-info">
						<a href="https://www.douban.com/people/212097281/" class="">bananana0</a>
						<div class="">
							1分钟前
							看过
							
						</div>
					</div>
				</li>
		</ul>
	
		
		<div class="subject-others-interests-ft">
			
				<a href="https://movie.douban.com/subject/25827935/collections">919862人看过</a>
					&nbsp;/&nbsp;
				<a href="https://movie.douban.com/subject/25827935/wishes">38401人想看</a>
		</div>
	
	</div>
	
	
	
		
		
	
	<!-- douban ad begin -->
	<div id="dale_movie_subject_middle_right"></div>
	<script type="text/javascript">
		(function (global) {
			if(!document.getElementsByClassName) {
				document.getElementsByClassName = function(className) {
					return this.querySelectorAll("." + className);
				};
				Element.prototype.getElementsByClassName = document.getElementsByClassName;
	
			}
			var articles = global.document.getElementsByClassName('article'),
				asides = global.document.getElementsByClassName('aside');
	
			if (articles.length > 0 && asides.length > 0 && articles[0].offsetHeight >= asides[0].offsetHeight) {
				(global.DoubanAdSlots = global.DoubanAdSlots || []).push('dale_movie_subject_middle_right');
			}
		})(this);
	</script>
	<!-- douban ad end -->
	
	
	
		<br/>
	
		
	<p class="pl">订阅七月与安生的评论: <br/><span class="feed">
		<a href="https://movie.douban.com/feed/subject/25827935/reviews"> feed: rss 2.0</a></span></p>
	
	
				</div>
				<div class="extra">
					
		
	<!-- douban ad begin -->
	<div id="dale_movie_subject_bottom_super_banner"></div>
	<script type="text/javascript">
		(function (global) {
			var body = global.document.body,
				html = global.document.documentElement;
	
			var height = Math.max(body.scrollHeight, body.offsetHeight, html.clientHeight, html.scrollHeight, html.offsetHeight);
			if (height >= 2000) {
				(global.DoubanAdSlots = global.DoubanAdSlots || []).push('dale_movie_subject_bottom_super_banner');
			}
		})(this);
	</script>
	<!-- douban ad end -->
	
	
				</div>
			</div>
		</div>
	
			
		<div id="footer">
				<div class="footer-extra"></div>
			
	<span id="icp" class="fleft gray-link">
		&copy; 2005－2020 douban.com, all rights reserved 北京豆网科技有限公司
	</span>
	
	<a href="https://www.douban.com/hnypt/variformcyst.py" style="display: none;"></a>
	
	<span class="fright">
		<a href="https://www.douban.com/about">关于豆瓣</a>
		· <a href="https://www.douban.com/jobs">在豆瓣工作</a>
		· <a href="https://www.douban.com/about?topic=contactus">联系我们</a>
		· <a href="https://www.douban.com/about/legal">法律声明</a>
		
		· <a href="https://help.douban.com/?app=movie" target="_blank">帮助中心</a>
		· <a href="https://www.douban.com/doubanapp/">移动应用</a>
		· <a href="https://www.douban.com/partner/">豆瓣广告</a>
	</span>
	
		</div>
	
		</div>
		<script type="text/javascript" src="https://img3.doubanio.com/misc/mixed_static/514e57f32902216.js"></script><script type="text/javascript">
			if (!Do.ready) {
				!function(){var t,e,n=document,r=window,i=window.__external_files_loaded=window.__external_files_loaded||{},o=window.__external_files_loading=window.__external_files_loading||{},a=function(t){return t.constructor===Array},s={autoLoad:!0,coreLib:["//img3.doubanio.com/js/jquery.min.js"],mods:{}},c=n.getElementsByTagName("script"),d=c[c.length-1],u=[],l=!1,f=[],h=function(t,e,r,a,s){var d=c[0];if(t){if(i[t])return o[t]=!1,void(a&&a(t,s));if(o[t])return void setTimeout(function(){h(t,e,r,a,s)},1);o[t]=!0;var u,l=e||t.toLowerCase().split(/\./).pop().replace(/[\?#].*/,"");if("js"===l?(u=n.createElement("script"),u.setAttribute("type","text/javascript"),u.setAttribute("src",t),u.setAttribute("async",!0)):"css"===l&&(u=n.createElement("link"),u.setAttribute("type","text/css"),u.setAttribute("rel","stylesheet"),u.setAttribute("href",t),i[t]=!0),u){if(r&&(u.charset=r),"css"===l)return d.parentNode.insertBefore(u,d),void(a&&a(t,s));u.onload=u.onreadystatechange=function(){this.readyState&&"loaded"!==this.readyState&&"complete"!==this.readyState||(i[this.getAttribute("src")]=!0,a&&a(this.getAttribute("src"),s),u.onload=u.onreadystatechange=null)},d.parentNode.insertBefore(u,d)}}},p=function(t){if(t&&a(t)){for(var e,n=0,r=[],i=s.mods,o=[],c={},d=function(t){var e,n,r=0;if(c[t])return o;if(c[t]=!0,i[t].requires){for(n=i[t].requires;"undefined"!=typeof(e=n[r++]);)i[e]?(d(e),o.push(e)):o.push(e);return o}return o};"undefined"!=typeof(e=t[n++]);)i[e]&&i[e].requires&&i[e].requires[0]&&(o=[],c={},r=r.concat(d(e))),r.push(e);return r}},y=function(){l=!0,u.length>0&&(e.apply(this,u),u=[])},m=function(){n.addEventListener?n.removeEventListener("DOMContentLoaded",m,!1):n.attachEvent&&n.detachEvent("onreadystatechange",m),y()},v=function(){if(!l){try{n.documentElement.doScroll("left")}catch(t){return r.setTimeout(v,1)}y()}},g=function(){if("complete"===n.readyState)return r.setTimeout(y,1);var t=!1;if(n.addEventListener)n.addEventListener("DOMContentLoaded",m,!1),r.addEventListener("load",y,!1);else if(n.attachEvent){n.attachEvent("onreadystatechange",m),r.attachEvent("onload",y);try{t=null===r.frameElement}catch(t){}document.documentElement.doScroll&&t&&v()}},E=function(t){t&&a(t)&&(this.queue=t,this.current=null)};E.prototype={_interval:10,start:function(){return this.current=this.next(),this.current?void this.run():void(this.end=!0)},run:function(){var t,e=this,n=this.current;return"function"==typeof n?(n(),void this.start()):void("string"==typeof n&&(s.mods[n]?(t=s.mods[n],h(t.path,t.type,t.charset,function(t){e.start()},e)):/\.js|\.css/i.test(n)?h(n,"","",function(t,e){e.start()},e):this.start()))},next:function(){return this.queue.shift()}},t=d.getAttribute("data-cfg-autoload"),"string"==typeof t&&(s.autoLoad="true"===t.toLowerCase()),t=d.getAttribute("data-cfg-corelib"),"string"==typeof t&&(s.coreLib=t.split(",")),e=function(){var t,e=[].slice.call(arguments);f.length>0&&(e=f.concat(e)),s.autoLoad&&(e=s.coreLib.concat(e)),t=new E(p(e)),t.start()},e.add=function(t,e){t&&e&&e.path&&(s.mods[t]=e)},e.delay=function(){var t=[].slice.call(arguments),n=t.shift();r.setTimeout(function(){e.apply(this,t)},n)},e.global=function(){var t=[].slice.call(arguments);f=f.concat(t)},e.ready=function(){var t=[].slice.call(arguments);return l?e.apply(this,t):void(u=u.concat(t))},e.css=function(t){var e=n.getElementById("do-inline-css");e||(e=n.createElement("style"),e.type="text/css",e.id="do-inline-css",n.getElementsByTagName("head")[0].appendChild(e)),e.styleSheet?e.styleSheet.cssText=e.styleSheet.cssText+t:e.appendChild(n.createTextNode(t))},s.autoLoad&&e(s.coreLib),e.define=e.add,e._config=s,e._mods=s.mods,this.Do=e,g()}();
	
			}
			Do.ready(
				'https://img3.doubanio.com/f/movie/b2a06a0332fc1526f4caaf8c76c2717d24da408d/js/movie/lib/jsrender.min.js',
				function(){
					$(document).on('click', '.cross span', function(e) {
						var $this = $(this);
						var $dialog = $this.parents('#tv-play-source');
						$dialog.remove();
					});
					$('body').bind('click', function(e) {
						var $this = $(e.target),
							$source = $('.play-source');
						if (!$this.is('.playBtn') && !$this.parents('.play-source').length) {
							$source.remove();
						}
					});
					var sources = {};
					sources[1] = [
								{play_link: "https://www.douban.com/link2/?url=http%3A%2F%2Fv.qq.com%2Fx%2Fcover%2Fq8742fxhp6pj7zv.html%3Fptag%3Ddouban.movie&amp;subtype=1&amp;type=online-video", ep: "1"},
					];
					sources[9] = [
								{play_link: "https://www.douban.com/link2/?url=http%3A%2F%2Fwww.iqiyi.com%2Fv_19rr9nifyg.html%3Fvfm%3Dm_331_dbdy&amp;subtype=9&amp;type=online-video", ep: "1"},
					];
					sources[3] = [
								{play_link: "https://www.douban.com/link2/?url=http%3A%2F%2Fv.youku.com%2Fv_show%2Fid_XMTc3NTgxOTIxNg%3D%3D.html%3Ftpa%3DdW5pb25faWQ9MzAwMDA4XzEwMDAwMl8wMl8wMQ%26refer%3Desfhz_operation.xuka.xj_00003036_000000_FNZfau_19010900&amp;subtype=3&amp;type=online-video", ep: "1"},
					];
					sources[7] = [
								{play_link: "https://www.douban.com/link2/?url=http%3A%2F%2Fwww.mgtv.com%2Fb%2F297788%2F3660147.html%3Fcxid%3D94l62sj44&amp;subtype=7&amp;type=online-video", ep: "1"},
					];
					sources[17] = [
								{play_link: "https://www.douban.com/link2/?url=https%3A%2F%2Fwww.ixigua.com%2Fcinema%2Falbum%2F6569374011929657864%3F%26utm_source%3Ddouban&amp;subtype=17&amp;type=online-video&amp;link2key=69c211df49", ep: "1"},
					];
					$('.playBtn').click(function(e){
						$('.play-source').remove();
	
						var height, width, tmpl, cn, source;
						var $dialog = $('#tv-play-source');
						var $this = $(this);
						source = $this.data('source');
						if(!source)return;
						cn = $this.data('cn');
	
						tmpl = $.templates('#episode-tmpl');
						$dialog = $(tmpl.render({playlist: sources[source], cn: cn}));
	
						$dialog.hide();
						$('body').append($dialog);
						width = $dialog.outerWidth();
	
						$dialog.css({
							marginLeft: -width / 2,
							top: $this.offset().top + $this.outerHeight()
						}).show();
					});
				});
		</script><script type="text/javascript" src="https://img3.doubanio.com/misc/mixed_static/53c8a070020f84ba.js"></script>
			
			
		<link rel="stylesheet" type="text/css" href="https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css" />
		<link rel="stylesheet" type="text/css" href="https://img3.doubanio.com/f/movie/4aca95d66d37ec0712b3d19973b5d8feb75f2f05/css/movie/mod/reg_login_pop.css" />
		<script type="text/javascript" src="https://img3.doubanio.com/f/shire/77323ae72a612bba8b65f845491513ff3329b1bb/js/do.js" data-cfg-autoload="false"></script>
		<script type="text/javascript" src="https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js"></script>
		<script type="text/javascript">
			var HTTPS_DB='https://www.douban.com';
	var account_pop={open:function(o,e){e?referrer="?referrer="+encodeURIComponent(e):referrer="?referrer="+window.location.href;var n="",i="",t=448;n="用户登录",i="https://accounts.douban.com/passport/login_popup?source=movie";var r=document.location.protocol+"//"+document.location.hostname,a=dui.Dialog({width:340,title:n,height:t,cls:"account_pop",isHideTitle:!0,modal:!0,content:"<iframe scrolling='no' frameborder='0' width='340' height='"+t+"' src='"+i+"' name='"+r+"'></iframe>"},!0),c=a.node;if(c.undelegate(),c.delegate(".dui-dialog-close","click",function(){var o=$("body");o.find("#login_msk").hide(),o.find(".account_pop").remove()}),$(window).width()<478){var d="";"reg"===o?d=HTTPS_DB+"/accounts/register"+referrer:"login"===o&&(d=HTTPS_DB+"/accounts/login"+referrer),window.location.href=d}else a.open();$(window).bind("message",function(o){"https://accounts.douban.com"===o.originalEvent.origin&&(c.find("iframe").css("height",o.originalEvent.data),c.height(o.originalEvent.data),a.update())})}};Douban&&Douban.init_show_login&&(Douban.init_show_login=function(o){var e=$(o);e.click(function(){var o=e.data("ref")||"";return account_pop.open("login",o),!1})}),Do(function(){$("body").delegate(".pop_register","click",function(o){o.preventDefault();var e=$(this).data("ref")||"";return account_pop.open("reg",e),!1}),$("body").delegate(".pop_login","click",function(o){o.preventDefault();var e=$(this).data("ref")||"";return account_pop.open("login",e),!1})});
		</script>
	
		
		
		
		
	
	
	
	
		
	<script type="text/javascript">
		(function (global) {
			var newNode = global.document.createElement('script'),
				existingNode = global.document.getElementsByTagName('script')[0],
				adSource = '//erebor.douban.com/',
				userId = '',
				browserId = '5TPL2nlZl7A',
				criteria = '7:马思纯|7:蔡纲|7:姚欣言|7:中国大陆|7:感动|7:剧情|7:周冬雨|7:友情|7:蒋亭轩|7:沙全泽|7:曾国祥|7:2016|7:成长|7:女性|7:小说改编|7:爱情|7:任可|7:蒙亭宜|7:文艺|7:李萍|7:青春|7:李程彬|7:陆忠|7:李昊芳|3:/subject/25827935/',
				preview = '',
				debug = false,
				adSlots = ['dale_movie_subject_top_icon', 'dale_movie_subject_top_right', 'dale_movie_subject_top_middle', 'dale_movie_subject_inner_middle', 'dale_movie_subject_download_middle', 'dale_movie_subject_banner_after_intro', 'dale_movie_review_right_guess_you_like'];
	
			global.DoubanAdRequest = {src: adSource, uid: userId, bid: browserId, crtr: criteria, prv: preview, debug: debug};
			global.DoubanAdSlots = (global.DoubanAdSlots || []).concat(adSlots);
	
			newNode.setAttribute('type', 'text/javascript');
			newNode.setAttribute('src', '//img1.doubanio.com/eDRjYjNvdi9mL2FkanMvZTQ2YTNkMjgwYjBiMzc2OWE4YTI3MWFhMzI0NTQwMTBlMWY3OTYzMy9hZC5yZWxlYXNlLmpz');
			newNode.setAttribute('async', true);
			existingNode.parentNode.insertBefore(newNode, existingNode);
		})(this);
	</script>
	
	
	
	
	
	
	
	
	
	
	
		
	  
	
	
	
	
	
	
	
	
	
	<script type="text/javascript">
	var _paq = _paq || [];
	_paq.push(['trackPageView']);
	_paq.push(['enableLinkTracking']);
	(function() {
		var p=(('https:' == document.location.protocol) ? 'https' : 'http'), u=p+'://fundin.douban.com/';
		_paq.push(['setTrackerUrl', u+'piwik']);
		_paq.push(['setSiteId', '100001']);
		var d=document, g=d.createElement('script'), s=d.getElementsByTagName('script')[0];
		g.type='text/javascript';
		g.defer=true;
		g.async=true;
		g.src=p+'://img3.doubanio.com/dae/fundin/piwik.js';
		s.parentNode.insertBefore(g,s);
	})();
	</script>
	
	<script type="text/javascript">
	var setMethodWithNs = function(namespace) {
	  var ns = namespace ? namespace + '.' : ''
		, fn = function(string) {
			if(!ns) {return string}
			return ns + string
		  }
	  return fn
	}
	
	var gaWithNamespace = function(fn, namespace) {
	  var method = setMethodWithNs(namespace)
	  fn.call(this, method)
	}
	
	var _gaq = _gaq || []
	  , accounts = [
		  { id: 'UA-7019765-1', namespace: 'douban' }
		, { id: 'UA-7019765-19', namespace: '' }
		]
	  , gaInit = function(account) {
		  gaWithNamespace(function(method) {
			gaInitFn.call(this, method, account)
		  }, account.namespace)
		}
	  , gaInitFn = function(method, account) {
		  _gaq.push([method('_setAccount'), account.id]);
		  _gaq.push([method('_setSampleRate'), '5']);
	
		  
	  _gaq.push([method('_addOrganic'), 'google', 'q'])
	  _gaq.push([method('_addOrganic'), 'baidu', 'wd'])
	  _gaq.push([method('_addOrganic'), 'soso', 'w'])
	  _gaq.push([method('_addOrganic'), 'youdao', 'q'])
	  _gaq.push([method('_addOrganic'), 'so.360.cn', 'q'])
	  _gaq.push([method('_addOrganic'), 'sogou', 'query'])
	  if (account.namespace) {
		_gaq.push([method('_addIgnoredOrganic'), '豆瓣'])
		_gaq.push([method('_addIgnoredOrganic'), 'douban'])
		_gaq.push([method('_addIgnoredOrganic'), '豆瓣网'])
		_gaq.push([method('_addIgnoredOrganic'), 'www.douban.com'])
	  }
	
		  if (account.namespace === 'douban') {
			_gaq.push([method('_setDomainName'), '.douban.com'])
		  }
	
			_gaq.push([method('_setCustomVar'), 1, 'responsive_view_mode', 'desktop', 3])
	
			_gaq.push([method('_setCustomVar'), 2, 'login_status', '0', 2]);
	
		  _gaq.push([method('_trackPageview')])
		}
	
	for(var i = 0, l = accounts.length; i < l; i++) {
	  var account = accounts[i]
	  gaInit(account)
	}
	
	
	;(function() {
		var ga = document.createElement('script');
		ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
		ga.setAttribute('async', 'true');
		document.documentElement.firstChild.appendChild(ga);
	})()
	</script>
	
	
	
	
	
	
	
	
		  
		
	
		<!-- dae-web-movie--default-7f4d549c79-jlh62-->
	
	  <script>_SPLITTEST=''</script>
	</body>
	
	</html>
	
	
	`

	c.Ctx.WriteString(models.GetMovieDirector(sMoviehtml) + "|")
	c.Ctx.WriteString(models.GetMovieName(sMoviehtml) + "|")
	c.Ctx.WriteString(models.GetMovieMainCharactor(sMoviehtml))

}
