package aklapi

const testHTML = `
<!DOCTYPE html >
<html lang="en" dir="ltr" class="ms-isBot">
    <head><!-- Version: 1.2.20030.01 -->
<!-- Google Tag Manager -->
<script>(function(w, d, s, l, i)
        {
            w[l] = w[l] ||[]; w[l].push({
                'gtm.start':
new Date().getTime(),event:'gtm.js'});var f = d.getElementsByTagName(s)[0],
j = d.createElement(s), dl = l != 'dataLayer' ? '&l=' + l : ''; j.async=true;j.src=
'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
})(window,document,'script','dataLayer','GTM-TB3J9W');</script>
<!-- End Google Tag Manager --><meta http-equiv="X-UA-Compatible" content="IE=EmulateIE10" /><meta http-equiv="Content-type" content="text/html; charset=utf-8" /><meta http-equiv="Expires" content="0" /><meta charset="utf-8" /><meta name="viewport" content="width=device-width" /><meta name="author" content="Auckland Council" /><meta name="format-detection" content="telephone=no" /><meta property="og:type" content="website"><meta property="og:site_name" content="Auckland Council"><meta property="og:title" content="Your collection day"><meta name="twitter:title" content="Your collection day"><meta property="og:description" content="Your collection day"><meta name="twitter:description" content="Your collection day"><meta property="og:url" content="https%3a%2f%2fwww.aucklandcouncil.govt.nz%2frubbish-recycling%2frubbish-recycling-collections%2fPages%2fcollection-day-detail.aspx%3fan%3d12341573910"><meta property="og:image" content="https://www.aucklandcouncil.govt.nz/SiteAssets/AC_pohutukawa-icon.png"><meta name="twitter:image" content="https://www.aucklandcouncil.govt.nz/SiteAssets/AC_pohutukawa-icon.png"><link rel="shortcut icon" href="/_layouts/15/ACWeb/images/favicon.ico" type="image/vnd.microsoft.icon" id="favicon" /><title>
	
                Your collection day 
            
</title><link rel="stylesheet" type="text/css" href="/_layouts/15/ACWeb/css/style.css?rev=MS4yLjIwMDMwLjAx"/>
<link rel="stylesheet" type="text/css" href="/_layouts/15/ACWeb/css/ACWeb.min.css?rev=MS4yLjIwMDMwLjAx"/>
<script type="text/javascript" src="/_layouts/15/init.js?rev=R500FBwHjH3TO5LeC6KbFQ%3D%3D"></script>
<script type="text/javascript" src="/_layouts/15/acweb/js/ribbonactions.min.js"></script>
<script type="text/javascript" src="/ScriptResource.axd?d=7W5GzupJRc8ieAvealpPKUPNUsmy748bS2HNOftrf-PygBU0wVMjFNbDzqr6uw1R-zuti0m_P4XHr_DiiaoSI-f4N71tPqO-5W6GGJ0mN7qhcMwoakbzoeN7tbY66fF9xKJGs-4PekQ_1UnR-6NOz4W4uVIzI7sXEoyo4Q_KCZppvrCsnXLGDR7J76bWsg080&amp;t=ffffffffecf19baa"></script>
<script type="text/javascript" src="/_layouts/15/blank.js?rev=ZaOXZEobVwykPO9g8hq%2F8A%3D%3D"></script>
<script type="text/javascript">RegisterSod("initstrings.js", "\u002f_layouts\u002f15\u002f1033\u002finitstrings.js?rev=S11vfGURQYVuACMEY0tLTg\u00253D\u00253D");</script>
<script type="text/javascript">RegisterSod("strings.js", "\u002f_layouts\u002f15\u002f1033\u002fstrings.js?rev=xXYZY4hciX287lShPZuClw\u00253D\u00253D");RegisterSodDep("strings.js", "initstrings.js");</script>
<script type="text/javascript">RegisterSod("sp.init.js", "\u002f_layouts\u002f15\u002fsp.init.js?rev=jvJC3Kl5gbORaLtf7kxULQ\u00253D\u00253D");</script>
<script type="text/javascript">RegisterSod("sp.res.resx", "\u002f_layouts\u002f15\u002fScriptResx.ashx?culture=en\u00252Dus\u0026name=SP\u00252ERes\u0026rev=yNk\u00252FhRzgBn40LJVP\u00252BqfgdQ\u00253D\u00253D");</script>
<script type="text/javascript">RegisterSod("sp.ui.dialog.js", "\u002f_layouts\u002f15\u002fsp.ui.dialog.js?rev=3Oh2QbaaiXSb7ldu2zd6QQ\u00253D\u00253D");RegisterSodDep("sp.ui.dialog.js", "sp.init.js");RegisterSodDep("sp.ui.dialog.js", "sp.res.resx");</script>
<script type="text/javascript">RegisterSod("core.js", "\u002f_layouts\u002f15\u002fcore.js?rev=GpU7vxyOqzS0F9OfEX3CCw\u00253D\u00253D");RegisterSodDep("core.js", "strings.js");</script>
<script type="text/javascript">RegisterSod("menu.js", "\u002f_layouts\u002f15\u002fmenu.js?rev=cXv35JACAh0ZCqUwKU592w\u00253D\u00253D");</script>
<script type="text/javascript">RegisterSod("mQuery.js", "\u002f_layouts\u002f15\u002fmquery.js?rev=VYAJYBo5H8I3gVSL3MzD6A\u00253D\u00253D");</script>
<script type="text/javascript">RegisterSod("callout.js", "\u002f_layouts\u002f15\u002fcallout.js?rev=ryx2n4ePkYj1\u00252FALmcsXZfA\u00253D\u00253D");RegisterSodDep("callout.js", "strings.js");RegisterSodDep("callout.js", "mQuery.js");RegisterSodDep("callout.js", "core.js");</script>
<script type="text/javascript">RegisterSod("clienttemplates.js", "\u002f_layouts\u002f15\u002fclienttemplates.js?rev=NjHQZTfQs\u00252BiDIjz4PK\u00252FeWg\u00253D\u00253D");RegisterSodDep("clienttemplates.js", "initstrings.js");</script>
<script type="text/javascript">RegisterSod("sharing.js", "\u002f_layouts\u002f15\u002fsharing.js?rev=XxxHIxIIc8BsW9ikVc6dgA\u00253D\u00253D");RegisterSodDep("sharing.js", "strings.js");RegisterSodDep("sharing.js", "mQuery.js");RegisterSodDep("sharing.js", "clienttemplates.js");RegisterSodDep("sharing.js", "core.js");</script>
<script type="text/javascript">RegisterSod("suitelinks.js", "\u002f_layouts\u002f15\u002fsuitelinks.js?rev=REwVU5jSsadDdOZlCx4wpA\u00253D\u00253D");RegisterSodDep("suitelinks.js", "strings.js");RegisterSodDep("suitelinks.js", "core.js");</script>
<script type="text/javascript">RegisterSod("sp.runtime.js", "\u002f_layouts\u002f15\u002fsp.runtime.js?rev=5f2WkYJoaxlIRdwUeg4WEg\u00253D\u00253D");RegisterSodDep("sp.runtime.js", "sp.res.resx");</script>
<script type="text/javascript">RegisterSod("sp.js", "\u002f_layouts\u002f15\u002fsp.js?rev=BJ9b7Ak8FOS3NwGiaTXOjA\u00253D\u00253D");RegisterSodDep("sp.js", "sp.runtime.js");RegisterSodDep("sp.js", "sp.ui.dialog.js");RegisterSodDep("sp.js", "sp.res.resx");</script>
<script type="text/javascript">RegisterSod("userprofile", "\u002f_layouts\u002f15\u002fsp.userprofiles.js?rev=p5tCOm\u00252FlHUwcfll7W3pKNw\u00253D\u00253D");RegisterSodDep("userprofile", "sp.runtime.js");</script>
<script type="text/javascript">RegisterSod("followingcommon.js", "\u002f_layouts\u002f15\u002ffollowingcommon.js?rev=jWqEDmcjCSPmnQw2ZIfItQ\u00253D\u00253D");RegisterSodDep("followingcommon.js", "strings.js");RegisterSodDep("followingcommon.js", "sp.js");RegisterSodDep("followingcommon.js", "userprofile");RegisterSodDep("followingcommon.js", "core.js");RegisterSodDep("followingcommon.js", "mQuery.js");</script>
<script type="text/javascript">RegisterSod("profilebrowserscriptres.resx", "\u002f_layouts\u002f15\u002fScriptResx.ashx?culture=en\u00252Dus\u0026name=ProfileBrowserScriptRes\u0026rev=J5HzNnB\u00252FO1Id\u00252FGI18rpRcw\u00253D\u00253D");</script>
<script type="text/javascript">RegisterSod("sp.ui.mysitecommon.js", "\u002f_layouts\u002f15\u002fsp.ui.mysitecommon.js?rev=Ua8qmZSU9nyf53S7PEyJwQ\u00253D\u00253D");RegisterSodDep("sp.ui.mysitecommon.js", "sp.init.js");RegisterSodDep("sp.ui.mysitecommon.js", "sp.runtime.js");RegisterSodDep("sp.ui.mysitecommon.js", "userprofile");RegisterSodDep("sp.ui.mysitecommon.js", "profilebrowserscriptres.resx");</script>
<script type="text/javascript">RegisterSod("browserScript", "\u002f_layouts\u002f15\u002fnon_ie.js?rev=W2q45TO627Zi6ztdktTOtA\u00253D\u00253D");RegisterSodDep("browserScript", "strings.js");</script>
<script type="text/javascript">RegisterSod("inplview", "\u002f_layouts\u002f15\u002finplview.js?rev=iMf5THfqukSYut7sl9HwUg\u00253D\u00253D");RegisterSodDep("inplview", "strings.js");RegisterSodDep("inplview", "core.js");RegisterSodDep("inplview", "sp.js");</script>
<link type="text/xml" rel="alternate" href="/rubbish-recycling/rubbish-recycling-collections/_vti_bin/spsdisco.aspx" />
            
            
            
            <meta name="robots" content="noindex" /><link rel="canonical" href="https://www.aucklandcouncil.govt.nz:443/rubbish-recycling/rubbish-recycling-collections/Pages/collection-day-detail.aspx?an=12341573910" /><meta name="keywords" content="Your collection day" /><meta name="description" content="Your collection day" />
            
            
        
        <!-- Google Ubuntu fonts -->
        <link href="https://fonts.googleapis.com/css?family=Ubuntu:300,300i&amp;subset=latin-ext" rel="stylesheet" /><link href="https://fonts.googleapis.com/css?family=Ubuntu:500,500i&amp;subset=latin-ext" rel="stylesheet" /><link href="https://fonts.googleapis.com/css?family=Ubuntu:700,700i&amp;subset=latin-ext" rel="stylesheet" />
        <!--[if IE 9]><link rel="stylesheet" type="text/css" href="/_layouts/15/ACWeb/css/ie9.css" /><![endif]--> 
        <link rel="stylesheet" type="text/css" href="/_layouts/15/ACWeb/css/ac-print.min.css" media="print" /></head>
    <body onhashchange="if (typeof(_spBodyOnHashChange) != 'undefined') _spBodyOnHashChange();">
        
<!-- Google Tag Manager (noscript) -->
<noscript><iframe src='https://www.googletagmanager.com/ns.html?id=GTM-TB3J9W'
height='0' width='0' style='display:none;visibility:hidden'></iframe></noscript>
<!-- End Google Tag Manager(noscript) -->
        
        <div id="imgPrefetch" style="display:none">
<img src="/_layouts/15/ACWeb/images/favicon.ico?rev=23" />
</div>

        <form method="post" action="./collection-day-detail.aspx?an=12341573910" id="aspnetForm">
<input type="hidden" name="_wpcmWpid" id="_wpcmWpid" value="" />
<input type="hidden" name="wpcmVal" id="wpcmVal" value="" />
<input type="hidden" name="MSOWebPartPage_PostbackSource" id="MSOWebPartPage_PostbackSource" value="" />
<input type="hidden" name="MSOTlPn_SelectedWpId" id="MSOTlPn_SelectedWpId" value="" />
<input type="hidden" name="MSOTlPn_View" id="MSOTlPn_View" value="0" />
<input type="hidden" name="MSOTlPn_ShowSettings" id="MSOTlPn_ShowSettings" value="False" />
<input type="hidden" name="MSOGallery_SelectedLibrary" id="MSOGallery_SelectedLibrary" value="" />
<input type="hidden" name="MSOGallery_FilterString" id="MSOGallery_FilterString" value="" />
<input type="hidden" name="MSOTlPn_Button" id="MSOTlPn_Button" value="none" />
<input type="hidden" name="__REQUESTDIGEST" id="__REQUESTDIGEST" value="0xEC047E2B8F670BE261928BFA71453A6F8ACB916FFD1CE2CB1C617AB11CFD71F3938622B714B2DCB28DF9855BAD5556D063CC52CBAC9843851E374E1C3BDC0547,10 Feb 2020 09:54:21 -0000" />
<input type="hidden" name="MSOSPWebPartManager_DisplayModeName" id="MSOSPWebPartManager_DisplayModeName" value="Browse" />
<input type="hidden" name="MSOSPWebPartManager_ExitingDesignMode" id="MSOSPWebPartManager_ExitingDesignMode" value="false" />
<input type="hidden" name="MSOWebPartPage_Shared" id="MSOWebPartPage_Shared" value="" />
<input type="hidden" name="MSOLayout_LayoutChanges" id="MSOLayout_LayoutChanges" value="" />
<input type="hidden" name="MSOLayout_InDesignMode" id="MSOLayout_InDesignMode" value="" />
<input type="hidden" name="_wpSelected" id="_wpSelected" value="" />
<input type="hidden" name="_wzSelected" id="_wzSelected" value="" />
<input type="hidden" name="MSOSPWebPartManager_OldDisplayModeName" id="MSOSPWebPartManager_OldDisplayModeName" value="Browse" />
<input type="hidden" name="MSOSPWebPartManager_StartWebPartEditingName" id="MSOSPWebPartManager_StartWebPartEditingName" value="false" />
<input type="hidden" name="MSOSPWebPartManager_EndWebPartEditing" id="MSOSPWebPartManager_EndWebPartEditing" value="false" />
<input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="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" />


<script type="text/javascript">
//<![CDATA[
var MSOWebPartPageFormName = 'aspnetForm';
var g_presenceEnabled = true;
var g_wsaEnabled = false;
var g_wsaQoSEnabled = false;
var g_wsaQoSDataPoints = [];
var g_wsaLCID = 1033;
var g_wsaListTemplateId = 850;
var g_wsaSiteTemplateId = 'BLANKINTERNET#0';
var _fV4UI=true;var _spPageContextInfo = {webServerRelativeUrl: "\u002frubbish-recycling\u002frubbish-recycling-collections", webAbsoluteUrl: "https:\u002f\u002fwww.aucklandcouncil.govt.nz\u002frubbish-recycling\u002frubbish-recycling-collections", siteAbsoluteUrl: "https:\u002f\u002fwww.aucklandcouncil.govt.nz", serverRequestPath: "\u002frubbish-recycling\u002frubbish-recycling-collections\u002fPages\u002fcollection-day-detail.aspx", layoutsUrl: "_layouts\u002f15", webTitle: "Rubbish and recycling collections", webTemplate: "53", tenantAppVersion: "0", isAppWeb: false, Has2019Era: true, webLogoUrl: "\u002fSiteAssets\u002fAC_pohutukawa-icon.png", webLanguage: 1033, currentLanguage: 1033, currentUICultureName: "en-US", currentCultureName: "en-NZ", clientServerTimeDelta: new Date("2020-02-10T09:54:21.3791549Z") - new Date(), siteClientTag: "5451$$15.0.5127.1000", crossDomainPhotosEnabled:false, webUIVersion:15, webPermMasks:{High:0,Low:196705},pageListId:"{0cd09468-050e-4fb0-b89d-1cdcd5af4fc6}",pageItemId:9, pagePersonalizationScope:1, alertsEnabled:true, siteServerRelativeUrl: "\u002f", allowSilverlightPrompt:'True'};document.onreadystatechange=fnRemoveAllStatus; function fnRemoveAllStatus(){removeAllStatus(true)};var _spWebPartComponents = new Object();//]]>
</script>

<script src="/_layouts/15/blank.js?rev=ZaOXZEobVwykPO9g8hq%2F8A%3D%3D" type="text/javascript"></script>
<script type="text/javascript">
//<![CDATA[
if (Hcf) { Hcf.RibbonActions.init(); }(function(){

        if (typeof(_spBodyOnLoadFunctions) === 'undefined' || _spBodyOnLoadFunctions === null) {
            return;
        }
        _spBodyOnLoadFunctions.push(function() {

            if (typeof(SPClientTemplates) === 'undefined' || SPClientTemplates === null || (typeof(APD_InAssetPicker) === 'function' && APD_InAssetPicker())) {
                return;
            }

            var renderFollowFooter = function(renderCtx,  calloutActionMenu)
            {
                if (renderCtx.ListTemplateType == 700) 
                    myDocsActionsMenuPopulator(renderCtx, calloutActionMenu);
                else
                    CalloutOnPostRenderTemplate(renderCtx, calloutActionMenu);

                var listItem = renderCtx.CurrentItem;
                if (typeof(listItem) === 'undefined' || listItem === null) {
                    return;
                }
                if (listItem.FSObjType == 0) {
                    calloutActionMenu.addAction(new CalloutAction({
                        text: Strings.STS.L_CalloutFollowAction,
                        tooltip: Strings.STS.L_CalloutFollowAction_Tooltip,
                        onClickCallback: function (calloutActionClickEvent, calloutAction) {
                            var callout = GetCalloutFromRenderCtx(renderCtx);
                            if (!(typeof(callout) === 'undefined' || callout === null))
                                callout.close();
                            SP.SOD.executeFunc('followingcommon.js', 'FollowSelectedDocument', function() { FollowSelectedDocument(renderCtx); });
                        }
                    }));
                }
            };

            var registerOverride = function(id) {
                var followingOverridePostRenderCtx = {};
                followingOverridePostRenderCtx.BaseViewID = 'Callout';
                followingOverridePostRenderCtx.ListTemplateType = id;
                followingOverridePostRenderCtx.Templates = {};
                followingOverridePostRenderCtx.Templates.Footer = function(renderCtx) {
                    var  renderECB;
                    if (typeof(isSharedWithMeView) === 'undefined' || isSharedWithMeView === null) {
                        renderECB = true;
                    } else {
                        var viewCtx = getViewCtxFromCalloutCtx(renderCtx);
                        renderECB = !isSharedWithMeView(viewCtx);
                    }
                    return CalloutRenderFooterTemplate(renderCtx, renderFollowFooter, renderECB);
                };
                SPClientTemplates.TemplateManager.RegisterTemplateOverrides(followingOverridePostRenderCtx);
            }
            registerOverride(101);
            registerOverride(700);
        });
    })();(function(){

        if (typeof(_spBodyOnLoadFunctions) === 'undefined' || _spBodyOnLoadFunctions === null) {
            return;
        }
        _spBodyOnLoadFunctions.push(function() 
        {
          ExecuteOrDelayUntilScriptLoaded(
            function()
            {
              var pairs = SP.ScriptHelpers.getDocumentQueryPairs();
              var followDoc, itemId, listId, docName;
              for (var key in pairs)
              {
                if(key.toLowerCase() == 'followdocument') 
                  followDoc = pairs[key];
                else if(key.toLowerCase() == 'itemid') 
                  itemId = pairs[key];
                else if(key.toLowerCase() == 'listid') 
                  listId = pairs[key];
                else if(key.toLowerCase() == 'docname') 
                  docName = decodeURI(pairs[key]);
              } 

              if(followDoc != null && followDoc == '1' && listId!=null && itemId != null && docName != null)
              {
                SP.SOD.executeFunc('followingcommon.js', 'FollowDocumentFromEmail', function() 
                { 
                  FollowDocumentFromEmail(itemId, listId, docName);
                });
              }

            }, 'SP.init.js');

        });
    })();if (typeof(DeferWebFormInitCallback) == 'function') DeferWebFormInitCallback();//]]>
</script>

<input type="hidden" name="__VIEWSTATEGENERATOR" id="__VIEWSTATEGENERATOR" value="4EBEA030" />
<input type="hidden" name="__VIEWSTATEENCRYPTED" id="__VIEWSTATEENCRYPTED" value="" />
<input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="rAyAFGCYHtfnNXTWvETeouWoJzC0QolEIrwSgzEXL5tb/ULTB//iNVZIWYP0yMldO3RNAyT44U7HInqktva+LVdHx7fvIHv27HhjK6Zt6dkuSSxqkgBq3IMNjwQnHVOg9DjbaxI2g/EvFZUD98EYNBpT1eAx7ATPkj3csGIUFYI=" />
            <span id="DeltaSPWebPartManager">
                
            </span>
            <input type="hidden" name="ctl00$ScriptManager" id="ctl00_ScriptManager" />
<script type="text/javascript">
//<![CDATA[
Sys.Application.setServerId("ctl00_ScriptManager", "ctl00$ScriptManager");
Sys.Application._enableHistoryInScriptManager();
//]]>
</script>

            
            <div id="DeltaPageStatusBar">
	
                <div id="pageStatusBar"></div>
            
</div>
            
            <div id="s4-workspace" class="ac-workspace">
                <div id="s4-bodyContainer">
                    <a accesskey="[" class="skip-main" href="#main">Skip to main content</a>
                    
                        
                    
                    
                        

                    
                    
                        <div class="logo-mobile-container hidden-sm-up noindex"> 
   <a title="Auckland Council" class="logo-foreground" href="/">
      <span class="sr-only">Auckland Council</span></a></div>
<header class="top fixedsticky noindex top fixedsticky s4-notdlg">
     <div class="header-content noindex"><div class="logoPrint"> 
      <img alt="Auckland Council" src="/_layouts/15/ACWeb/images/ac-logo-large.svg" /> 
      <br>
   </div><div class="hidden-xs-down"> 
      <a title="Auckland Council" class="logo-foreground" accesskey="1" href="/">
         <span class="sr-only">Auckland Council</span></a> </div><div class="search-and-menu"><div class="header-inner-wrap"><div class="textbox-with-a-button master-search noindex"> 
            <label class="sr-only" id="master-search-label">Search Auckland Council</label> 
            <div class="inner noindex">
               <input id="master-search" aria-labelledby="master-search-label" type="search" placeholder="Search Auckland Council" /></div><div class="btn-wrap noindex">
               <button class="btn btn-primary icon-button inside-of-a-textbox" type="submit"><span class="sr-only">Search</span></button></div></div><div class="alert global-search-address" style="display&#58;none;z-index&#58;0;"><div class="heading"> 
               <strong class="info">Looking for information on an address?</strong> Use the search within the relevant topic below. 
               <button class="close" aria-label="Close" type="button" data-dismiss="alert"> 
                  <span aria-hidden="true">Ã—</span> </button> </div><ul class="list"><li> 
                  <a href="/plans-projects-policies-reports-bylaws/our-plans-strategies/unitary-plan/Pages/default.aspx">Unitary Plan</a> </li><li> 
                  <a href="/property-rates-valuations/Pages/find-property-rates-valuation.aspx">Find your property rates or valuation</a> </li><li> 
                  <a href="/buying-property/propertyinformationgeomaps/Pages/default.aspx">Property information in GeoMaps</a> </li><li> 
                  <a href="/buying-property/order-property-report/Pages/default.aspx">LIMS and property files</a> </li><li> 
                  <a href="/rubbish-recycling/rubbish-recycling-collections/Pages/default.aspx">Rubbish and recycling collections</a> </li></ul></div></div><div class="noindex" style="padding-right&#58;10px;padding-left&#58;5px;"> 
         <a class="btn btn-secondary btn-menu" role="button" aria-expanded="false" href="#"><span><span class="css-icon-menu"><span class="css-icon-menu-line ci-ml1"></span><span class="css-icon-menu-line ci-ml2"></span><span class="css-icon-menu-line ci-ml3"></span><span class="css-icon-menu-line-bag ci-mlxb"><span class="css-icon-menu-line ci-mlx"></span> </span> </span> Menu </span> </a> </div><div> 
         <a class="btn btn-primary icon-button icon-aligned-to-the-left myaccount hidden-sm-down" role="button" href="#"><span class="light">my</span>AUCKLAND 
            <span class="light hidden-md-down">login</span> </a> </div></div></div><div class="menu-mobile hidden-md-up">
   <a class="btn btn-primary icon-button icon-aligned-to-the-left pull-right myaccount hidden-md-up" role="button" href="#"> 
      <span class="light">my</span>AUCKLAND </a> </div>
    

<nav class="navigation-menu">
        <div class="container">
        <div class="row noindex">
            <div class="col-md-12 col-lg-3">
                <h2 class="inverse">Browse our website</h2>
            </div>
            
        <div class="col-md-4 col-lg-3">
            <ul class="list-unstyled spaced">
                
        <li><a href='/property-rates-valuations/Pages/default.aspx' class="border-link inverse bold"> Property rates and valuations</a></li>
    
        <li><a href='/buying-property/Pages/default.aspx' class="border-link inverse bold"> Buying property</a></li>
    
        <li><a href='/licences-regulations/Pages/default.aspx' class="border-link inverse bold"> Licences and regulations</a></li>
    
        <li><a href='/plans-projects-policies-reports-bylaws/Pages/default.aspx' class="border-link inverse bold"> Plans, policies, bylaws, reports and projects</a></li>
    
        <li><a href='/environment/Pages/default.aspx' class="border-link inverse bold"> Environment</a></li>
    
        <li><a href='/about-auckland-council/Pages/home.aspx' class="border-link inverse bold"> About Auckland Council</a></li>
    
            </ul>
        </div>
    
        <div class="col-md-4 col-lg-3">
            <ul class="list-unstyled spaced">
                
        <li><a href='/arts-culture-heritage/Pages/default.aspx' class="border-link inverse bold"> Arts, culture and heritage</a></li>
    
        <li><a href='/building-and-consents/Pages/default.aspx' class="border-link inverse bold"> Building and consents</a></li>
    
        <li><a href='/cemeteries/Pages/default.aspx' class="border-link inverse bold"> Cemeteries</a></li>
    
        <li><a href='/dogs-animals/Pages/default.aspx' class="border-link inverse bold"> Dogs and other animals</a></li>
    
        <li><a href='/have-your-say/Pages/home.aspx' class="border-link inverse bold"> Have your say and help shape Auckland</a></li>
    
        <li><a href='/parks-recreation/Pages/default.aspx' class="border-link inverse bold"> Parks, recreation and community venues</a></li>
    
            </ul>
        </div>
    
        <div class="col-md-4 col-lg-3">
            <ul class="list-unstyled spaced">
                
        <li><a href='/report-problem/Pages/default.aspx' class="border-link inverse bold"> Contact us or report a problem</a></li>
    
        <li><a href='/rubbish-recycling/Pages/default.aspx' class="border-link inverse bold"> Rubbish and recycling</a></li>
    
        <li><a href='/grants-community-support-housing/Pages/default.aspx' class="border-link inverse bold"> Grants, community support and housing</a></li>
    
        <li><a href='/mayor-of-auckland/Pages/default.aspx' class="border-link inverse bold"> Mayor of Auckland</a></li>
    
            </ul>
        </div>
    

        </div>
    </div>
    <button class="close-button" onclick="return false;"><span class="sr-only">Close</span></button>

    <div class="fade-out"></div>
</nav>

    <div class="alert-civildefence noindex">
        
    </div>
</header>



<div class="hidden-sm-down">
    <div id="ctl00_PlaceHolderHeader_ACWebHeader_ACHeroImage_HeroImageSide">
	
        <img id="ctl00_PlaceHolderHeader_ACWebHeader_ACHeroImage_HeroPicture" class="hero" alt="" src="/Lists/backgroundimages/Waitemata_1.jpg" />
    
</div>
</div>


                    
                    <div class="section-wrapper">
                        <div id="mainArea" class="row-fluid">
                            
                            <div id="ctl00_DeltaPlaceHolderMain">
                                <section>
                                    
                                    
                                    
                                    <script type="text/javascript">// <![CDATA[ 


                                        _spBodyOnLoadFunctionNames.push("setupPageDescriptionCallout");
                                    // ]]>
</script>
                                    
                                    <div id="ctl00_mainPanel" class="mainPanel">
                                        
                                            <ul id="ctl00_PlaceHolderTitleBreadcrumb_SimpleSiteMapPath1" class="breadcrumb-wrapper">
	<li class="border-link"><a class="border-link" href="/Pages/default.aspx">Home</a></li><li class="border-link"><a title="Collection dates, see what you can put in your rubbish, recycling or food scraps bins, report a missed collection or dumped rubbish, find ways to dispose of unwanted items, or book an inorganic collection." class="border-link" href="/rubbish-recycling/Pages/default.aspx">Rubbish and recycling</a></li><li class="border-link"><a title="Find out what you can put in your rubbish and recycling and when your collection days are." class="border-link" href="/rubbish-recycling/rubbish-recycling-collections/Pages/default.aspx">Rubbish and recycling collections</a></li><li>Your collection day</li>
</ul>
                                        
                                        
                                        
                                        
                                        
                                        
<div id="generalHeading">
	
    

</div>
	<div class="welcome blank-wp">
		
        
        
		<div class="welcome-content">
            <div id="ctl00_PlaceHolderMain_ctl02_label" style='display:none'>Page Content</div><div id="ctl00_PlaceHolderMain_ctl02__ControlWrapper_RichHtmlField" class="ms-rtestate-field" style="display:inline" aria-labelledby="ctl00_PlaceHolderMain_ctl02_label"></div>
        </div>		
        <div class="cell-margin">
            <div class="ms-webpart-zone ms-fullWidth">
	<div id="MSOZoneCell_WebPartctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d" class="s4-wpcell-plain ms-webpartzone-cell ms-webpart-cell-vertical ms-fullWidth ">
		<div class="ms-webpart-chrome ms-webpart-chrome-vertical ms-webpart-chrome-fullWidth ">
			<div WebPartID="dfe289d2-6a8a-414d-a384-fc25a0db9a6d" HasPers="false" id="WebPartctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d" width="100%" class="ms-WPBody " allowDelete="false" allowExport="false" style="" ><div id="ctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d">
				

<h1>Your collection day</h1>

<div id="ctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d_ctl00_pnlDetail">
					
<h2 class="heading-te-reo m-b-3dot5" lang="mi">Ngā kōrero kohinga mōu</h2>
<h2 class="m-b-2">Red Square, Moscow</h2>
    
    <p></p>
     
            <div class="card-content m-b-2">
                <div class="card-header">
                    <h3 class="card-title h2">Household collection</h3>
                    <h4 class="h6"><span class="icon-rubbish"><span class="sr-only">Rubbish</span></span> Rubbish</h4>
                    <p>Collection day: <strong>Tuesday, weekly</strong> except after a <a class='border-link' href='/rubbish-recycling/rubbish-recycling-collections/Pages/public-holiday-collections.aspx'>public holiday</a>.</p>
                    
                    <h4 class="h6 m-t-1"><span class="icon-recycle"><span class="sr-only">Recycle</span></span> Recycling</h4>
                    <p>Collection day: <strong>Tuesday, fortnightly</strong> except after a <a class='border-link' href='/rubbish-recycling/rubbish-recycling-collections/Pages/public-holiday-collections.aspx'>public holiday</a>.</p>
                    Put bins out the night before or before 7am.
                    
                </div>
                <div id="ctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d_ctl00_pnlHouseholdBlock" class="card-block">
						
                    <h4 class='h6 m-t-1' style='margin-left: 10px;'><strong>Your next collection dates:</strong></h4>
                    
                        <div class="links"><span class='m-r-1'>Tuesday 11 February</span><span class='icon-rubbish'><span class='sr-only'>Rubbish</span></span> <span class='icon-recycle'><span class='sr-only'>Recycle</span></span> </div>
                    
                        <div class="links"><span class='m-r-1'>Tuesday 18 February</span><span class='icon-rubbish'><span class='sr-only'>Rubbish</span></span> </div>
                    
                    <div class="links"><a id="ctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d_ctl00_lnkWhere" href="javascript:WebForm_DoPostBackWithOptions(new WebForm_PostBackOptions(&quot;ctl00$SPWebPartManager1$g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d$ctl00$lnkWhere&quot;, &quot;&quot;, true, &quot;&quot;, &quot;&quot;, false, true))"><span class="docs">Where you can put your rubbish and recycling for collection</span></a>  </div>
                    <div class="links"><a id="ctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d_ctl00_hlRecycle" href="/rubbish-recycling/rubbish-recycling-collections/Pages/what-put-your-recycling.aspx"><span class="docs">What you can put in your recycle bin</span></a></div>
                    <div class="links"><a id="ctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d_ctl00_hlRubbish" href="/rubbish-recycling/rubbish-recycling-collections/Pages/what-put-your-rubbish.aspx"><span class="docs">What you can put in your rubbish bin</span></a></div>              
                
					</div>
            </div>
    
    <div class='card card-shadow document'><a class='card-inner-wrap' href='/CollectionCalendar/waitakerecalendar.pdf'>
            <div class='card-block'><p class='h2 card-title m-b-0'><span class='download'>
            <span>Rubbish and recycling collection calendar</span></span></p></div><div class='card-footer'>
            <strong>PDF</strong> 420.3 KB</div></a></div>

    <div class="card-content m-b-2">
                <div class="card-header">
                    <h3 class="card-title h2">Commercial collection</h3>
                    <h4 class="h6"><span class="icon-rubbish"><span class="sr-only">Rubbish</span></span> Rubbish</h4>
                    <p>Collection days: <strong>Tuesday, Thursday and Saturday, weekly</strong>. See <a class='border-link' href='/rubbish-recycling/rubbish-recycling-collections/Pages/public-holiday-collections.aspx'>public holidays</a> for exceptions.</p>
                    
                    <h4 class="h6 m-t-1"><span class="icon-recycle"><span class="sr-only">Recycle</span></span> Recycling</h4>
                    <p>Collection day: <strong>Tuesday, fortnightly</strong> except after a <a class='border-link' href='/rubbish-recycling/rubbish-recycling-collections/Pages/public-holiday-collections.aspx'>public holiday</a>.</p>
                    Put bins out the night before or before 7am.
                    
                </div>
                <div id="ctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d_ctl00_pnlCommercialBlock" class="card-block">
						
                    <h4 class='h6 m-t-1' style='margin-left: 10px;'><strong>Your next collection dates:</strong></h4>
                    
                            <div class="links"><span class='m-r-1'>Tuesday 11 February</span><span class='icon-rubbish'><span class='sr-only'>Rubbish</span></span> <span class='icon-recycle'><span class='sr-only'>Recycle</span></span> </div>
                        
                            <div class="links"><span class='m-r-1'>Tuesday 18 February</span><span class='icon-rubbish'><span class='sr-only'>Rubbish</span></span> </div>
                        
                    <div class="links"><a id="ctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d_ctl00_lnkWhere2" href="javascript:WebForm_DoPostBackWithOptions(new WebForm_PostBackOptions(&quot;ctl00$SPWebPartManager1$g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d$ctl00$lnkWhere2&quot;, &quot;&quot;, true, &quot;&quot;, &quot;&quot;, false, true))"><span class="docs">Where you can put your rubbish and recycling for collection</span></a>  </div>
                    <div class="links"><a id="ctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d_ctl00_hlRecycle2" href="/rubbish-recycling/rubbish-recycling-collections/Pages/what-put-your-recycling.aspx"><span class="docs">What you can put in your recycle bin</span></a></div>
                    <div class="links"><a id="ctl00_SPWebPartManager1_g_dfe289d2_6a8a_414d_a384_fc25a0db9a6d_ctl00_hlRubbish2" href="/rubbish-recycling/rubbish-recycling-collections/Pages/what-put-your-rubbish.aspx"><span class="docs">What you can put in your rubbish bin</span></a></div>              
                
					</div>
   </div>

				</div>


			</div><div class="ms-clear"></div></div>
		</div>
	</div>
</div>
        </div>
        <div class="cell-margin">
            <menu class="ms-hide">
	<ie:menuitem id="MSOMenu_Help" iconsrc="/_layouts/15/images/HelpIcon.gif" onmenuclick="MSOWebPartPage_SetNewWindowLocation(MenuWebPart.getAttribute('helpLink'), MenuWebPart.getAttribute('helpMode'))" text="Help" type="option" style="display:none">

	</ie:menuitem>
</menu>
        </div>
        <div class="welcome-content">
            <div id="ctl00_PlaceHolderMain_ctl03_label" style='display:none'>Additional Content</div><div id="ctl00_PlaceHolderMain_ctl03__ControlWrapper_RichHtmlField" class="ms-rtestate-field" style="display:inline" aria-labelledby="ctl00_PlaceHolderMain_ctl03_label"></div>
        </div>
        <div class="cell-margin">
            
        </div>
	</div>
        

<div class="social-media-wrapper m-b-3">
    <div class="tooltip bs-tether-element-attached-bottom fade in" role="tooltip">
            <div class="tooltip-arrow"></div>
            <div class="tooltip-inner">
                <span>Share to: </span>
                
                        <a target="_blank" href="mailto:?Subject=Shared%20from%20Auckland%20Council&body=Your%20collection%20day%0D%0AYour%20collection%20day%0D%0Ahttps://www.aucklandcouncil.govt.nz:443/rubbish-recycling/rubbish-recycling-collections/Pages/collection-day-detail.aspx?an=12341573910" class="social social__Mail To" data-toggle="tooltip" data-placement="top" title="" data-original-title="Mail To"><span class="sr-only">Mail To</span>
                            <img src="https://www.aucklandcouncil.govt.nz/SiteAssets/icon-social-email.svg" alt="Mail To" />
                        </a>
                    
                        <a target="_blank" href="https://www.facebook.com/sharer/sharer.php?u=https://www.aucklandcouncil.govt.nz:443/rubbish-recycling/rubbish-recycling-collections/Pages/collection-day-detail.aspx?an=12341573910&title=Your%20collection%20day&description=Your%20collection%20day" class="social social__Facebook" data-toggle="tooltip" data-placement="top" title="" data-original-title="Facebook"><span class="sr-only">Facebook</span>
                            <img src="https://www.aucklandcouncil.govt.nz/SiteAssets/icon-social-facebook.svg" alt="Facebook" />
                        </a>
                    
                        <a target="_blank" href="https://www.linkedin.com/shareArticle?url=https://www.aucklandcouncil.govt.nz:443/rubbish-recycling/rubbish-recycling-collections/Pages/collection-day-detail.aspx?an=12341573910&title=Your%20collection%20day&summary=Your%20collection%20day&token=&isFramed=true" class="social social__LinkedIn" data-toggle="tooltip" data-placement="top" title="" data-original-title="LinkedIn"><span class="sr-only">LinkedIn</span>
                            <img src="https://www.aucklandcouncil.govt.nz/SiteAssets/icon-social-linkedin.svg" alt="LinkedIn" />
                        </a>
                    
                        <a target="_blank" href="https://twitter.com/intent/tweet?url=https://www.aucklandcouncil.govt.nz:443/rubbish-recycling/rubbish-recycling-collections/Pages/collection-day-detail.aspx?an=12341573910&via=aklcouncil&text=Your%20collection%20day" class="social social__Twitter" data-toggle="tooltip" data-placement="top" title="" data-original-title="Twitter"><span class="sr-only">Twitter</span>
                            <img src="https://www.aucklandcouncil.govt.nz/SiteAssets/icon-social-twitter.svg" alt="Twitter" />
                        </a>
                    
            </div>
    </div>
    <button onclick="javascript:return false;" class="btn btn-secondary icon-button icon-aligned-to-the-left share" data-toggle-share=""><span>Share this page</span></button>
</div>

    
<div class="top-margin-lg">
    <div id="FeedbackOptions" class="btn-icon-group feedback" >
        <p><span id="ctl00_PlaceHolderMain_ACFeedback_feedbackLeadText">Is the information on this page helpful?</span></p>
        <button class="btn yes" aria-label="Yes the page information is helpful" type="button"><span aria-hidden="true">Yes</span></button>
        <button class="btn no" aria-label="No the page information is not helpful" type="button"><span aria-hidden="true">No</span></button>
    </div>
    <div id="FeedbackPanel" style="display:none;">
        <div tabindex="0" class="search-heading" role="alert">
            <p class="h4">Thanks for your feedback</p>
            <p></p>
        </div>
        <p class="form-field--label">To ask for help or report a problem with our services or facilities, <a id="ctl00_PlaceHolderMain_ACFeedback_lnkContactUs" href="https://www.aucklandcouncil.govt.nz/report-problem/Pages/default.aspx">contact us</a>.</p>
        <label class="form-field--label" for="feedback-text">Tell us how we can improve the information on this page. <span class="hint">(optional)</span></label>
        <textarea id="feedback-text" rows="6"></textarea>
        <p class="feedback-buttons">
            <button class="btn btn-secondary" id="SkipCommentsButton" >Skip</button>
            <button class="btn btn-primary" id="SendCommentsButton" enableviewstate="false"><span>Send</span></button>
        </p>
    </div>
    <div id="FeedbackThanksPanel" style="display:none;">
        <div class="search-heading" tabindex="0" role="alert">
            <p class="h4" style="display: block !important">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" width="20" height="20" class="green-tick">
                    <path class="tick" d="M6.5,17.25a1.246,1.246,0,0,1-.884-.366l-5-5a1.25,1.25,0,0,1,1.768-1.768L6.5,14.232,17.616,3.116a1.25,1.25,0,0,1,1.768,1.768l-12,12A1.246,1.246,0,0,1,6.5,17.25Z"/></path></svg>
                Thanks for your feedback
            </p>
        </div>
    </div>
</div>
	
<div style='display:none' id='hidZone'></div>
                                        <div class="mainPanelPadding "></div>
                                    </div>
                                    
                                    
                                </section>
                            </div>
                        </div>
                    </div>
                    
                        <div class="footer-wrapper noindex s4-notdlg">
                            <footer><a class="btn btn-primary icon-button top" role="button" href="#">Top</a>
   <div class="content noindex"><h2 class="h1 centered inverse">Need help?<br></h2><p class="centered inverse">Get in touch with us.</p><div class="container-fluid m-x-minus15px"><div class="row"><div class="col-xl-4 centered m-b-1 noindex">
               <a class="btn btn-primary icon-button icon-aligned-to-the-left computer" role="button" accesskey="0" href="/report-problem/Pages/default.aspx">Contact us</a> </div><div class="col-xl-4 centered m-b-1 noindex">
               <a class="btn btn-primary icon-button icon-aligned-to-the-left phone" role="button" accesskey="9" href="/093010101"><span class="sr-only">Call us on </span>09 301 0101</a> </div><div class="col-xl-4 centered m-b-1 noindex">
               <a class="btn btn-primary icon-button icon-aligned-to-the-left person" role="button" href="/report-problem/visit-us/Pages/default.aspx">Visit us</a> </div></div></div><div class="separator">&#160;</div><h3 class="h2 inverse m-b-1">Other Auckland Council websites</h3><div class="m-t-1 footer-links"><div class="row"><div class="col-md-6"><ul class="list-unstyled spaced"><li> 
                     <a class="H3 border-link inverse external" href="http&#58;//www.aucklandnz.com/">Auckland Tourism, Events and Economic Development</a> </li><li> 
                     <a class="H3 border-link inverse external" href="https&#58;//at.govt.nz/">Auckland Transport</a> </li><li> 
                     <a class="H3 border-link inverse external" href="https&#58;//www.aucklandlibraries.govt.nz/">Libraries</a> </li></ul></div><div class="col-md-6"><ul class="list-unstyled spaced"><li> 
                     <a class="H3 border-link inverse external" href="https&#58;//www.panuku.co.nz/">
                        <span lang="mi">Panuku</span> Development Auckland</a> </li><li> 
                     <a class="H3 border-link inverse external" href="http&#58;//www.rfal.co.nz/">Regional Facilities Auckland</a> </li><li> 
                     <a class="H3 border-link inverse external" href="https&#58;//www.watercare.co.nz/">Watercare Services</a> </li><li> 
                     <a class="H3 border-link inverse" href="/Pages/auckland-council-websites.aspx">More Auckland Council websites…</a> </li></ul></div></div></div></div> </footer>
<div class="footer-more"><div class="social-buttons noindex"> 
      <a title="Auckland Council on Facebook" class="social social__facebook" href="https&#58;//www.facebook.com/aklcouncil">
         <span class="sr-only">Auckland Council on Facebook</span></a><a title="Auckland Council on Twitter" class="social social__twitter" href="https&#58;//twitter.com/aklcouncil"><span class="sr-only">Auckland Council on Twitter</span></a><a title="Auckland Council on LinkedIn" class="social social__linkedin" href="https&#58;//www.linkedin.com/company/auckland-council" target="_blank"><span class="sr-only">Auckland Council on LinkedIn</span></a><a title="Auckland Council on Youtube" class="social social__youtube" href="https&#58;//www.youtube.com/user/AklCouncil"><span class="sr-only">Auckland Council on Youtube</span></a> </div><div title="Auckland Council" class="logo">&#160;</div><div class="centered terms-and-conditions"> 
      <a class="border-link" href="/Pages/terms-and-conditions.aspx">Terms and conditions</a><a class="border-link" href="/Pages/privacy-policy.aspx">Privacy policy</a><a class="border-link" accesskey="2" href="/Pages/site-map.aspx">Site map</a><br> © Auckland Council 2019<br><a title="New Zealand Government" accesskey="/" href="https&#58;//www.govt.nz/" target="_blank"><img alt="New Zealand Government" src="/_layouts/15/acweb/images/NZGovt-logo.svg" /></a> </div></div>

                        </div>
                    
                </div>
            </div>

            <!-- ===== STARTER: Main Scrolling Body Ends Here ================================================================================= -->

            <!-- ===== STARTER: Needed for form stuff ========================================================================================= -->
            <div id="DeltaFormDigest">
	
                
                    <script type="text/javascript">//<![CDATA[
        var formDigestElement = document.getElementsByName('__REQUESTDIGEST')[0];
        if (!((formDigestElement == null) || (formDigestElement.tagName.toLowerCase() != 'input') || (formDigestElement.type.toLowerCase() != 'hidden') ||
            (formDigestElement.value == null) || (formDigestElement.value.length <= 0)))
        {
            formDigestElement.value = '0xEC047E2B8F670BE261928BFA71453A6F8ACB916FFD1CE2CB1C617AB11CFD71F3938622B714B2DCB28DF9855BAD5556D063CC52CBAC9843851E374E1C3BDC0547,10 Feb 2020 09:54:21 -0000';
            g_updateFormDigestPageLoaded = new Date();
        }
        //]]>
        </script>
                
            
</div>
        

<script type="text/javascript">
//<![CDATA[
var _spFormDigestRefreshInterval = 1440000;window.g_updateFormDigestPageLoaded = new Date(); window.g_updateFormDigestPageLoaded.setDate(window.g_updateFormDigestPageLoaded.getDate() -5);var _fV4UI = true;
function _RegisterWebPartPageCUI()
{
    var initInfo = {editable: false,isEditMode: false,allowWebPartAdder: false,listId: "{0cd09468-050e-4fb0-b89d-1cdcd5af4fc6}",itemId: 9,recycleBinEnabled: true,enableMinorVersioning: true,enableModeration: true,forceCheckout: true,rootFolderUrl: "\u002frubbish-recycling\u002frubbish-recycling-collections\u002fPages",itemPermissions:{High:0,Low:196705}};
    SP.Ribbon.WebPartComponent.registerWithPageManager(initInfo);
    var wpcomp = SP.Ribbon.WebPartComponent.get_instance();
    var hid;
    hid = document.getElementById("_wpSelected");
    if (hid != null)
    {
        var wpid = hid.value;
        if (wpid.length > 0)
        {
            var zc = document.getElementById(wpid);
            if (zc != null)
                wpcomp.selectWebPart(zc, false);
        }
    }
    hid = document.getElementById("_wzSelected");
    if (hid != null)
    {
        var wzid = hid.value;
        if (wzid.length > 0)
        {
            wpcomp.selectWebPartZone(null, wzid);
        }
    }
};
function __RegisterWebPartPageCUI() {
ExecuteOrDelayUntilScriptLoaded(_RegisterWebPartPageCUI, "sp.ribbon.js");}
_spBodyOnLoadFunctionNames.push("__RegisterWebPartPageCUI");var __wpmExportWarning='This Web Part Page has been personalized. As a result, one or more Web Part properties may contain confidential information. Make sure the properties contain information that is safe for others to read. After exporting this Web Part, view properties in the Web Part description file (.WebPart) by using a text editor such as Microsoft Notepad.';var __wpmCloseProviderWarning='You are about to close this Web Part.  It is currently providing data to other Web Parts, and these connections will be deleted if this Web Part is closed.  To close this Web Part, click OK.  To keep this Web Part, click Cancel.';var __wpmDeleteWarning='You are about to permanently delete this Web Part.  Are you sure you want to do this?  To delete this Web Part, click OK.  To keep this Web Part, click Cancel.';var g_clientIdDeltaPlaceHolderMain = "ctl00_DeltaPlaceHolderMain";
var g_clientIdDeltaPlaceHolderUtilityContent = "DeltaPlaceHolderUtilityContent";
//]]>
</script>
</form>

        <!-- ===== STARTER: Hidden Placeholders =========================================================================================== -->
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        

        <!-- Include all compiled plugins (below), or include individual files as needed -->
        <script type="text/javascript" src="/_layouts/15/ACWeb/js/jquery.min.js?rev=MS4yLjIwMDMwLjAx"></script>
        <script type="text/javascript" src="/_layouts/15/ACWeb/js/acscripts-vendor.min.js?rev=MS4yLjIwMDMwLjAx"></script>
        <script type="text/javascript" src="/_layouts/15/ACWeb/js/acscripts-main.min.js?rev=MS4yLjIwMDMwLjAx"></script>
        <!--<script type="text/javascript" src="/_layouts/15/ACWeb/js/scripts.min.js?rev=MS4yLjIwMDMwLjAx"></script>-->
        <script type="text/javascript" src="/_layouts/15/ACWeb/js/main.min.js?rev=MS4yLjIwMDMwLjAx"></script>
        <script type="text/javascript" src="/_layouts/15/ACWeb/js/AddressSearch.min.js?rev=MS4yLjIwMDMwLjAx"></script>
        
        

                <!-- ===== STARTER: Adds extra stuff like another form for Survey management ====================================================== -->
        <span id="DeltaPlaceHolderUtilityContent">
            
        </span>
    </body>
</html>
`
