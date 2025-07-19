package mocks

var OID_response = `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://www.w3.org/2003/05/soap-envelope">
<S:Body>
<ds:getUserOidUserShortnameResponse xmlns:ds="http://www.dotsource.de">
<Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1690198889228_JUser">
<Part name="lineNumber">1</Part>
<Part name="oid">1690198889228_JUser</Part>
<Part name="%loadername%">getUserOidUserShortnameLoader</Part>
</Record>
</ds:getUserOidUserShortnameResponse>
</S:Body>
</S:Envelope>`

var Reports_response = `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://www.w3.org/2003/05/soap-envelope">
    <S:Body>
        <ds:getEffortsByOidAndDateRangeResponse xmlns:ds="http://www.dotsource.de">
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1745863747437_JEffort">
                <Part name="effortExpense">165</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Academy</Part>
                <Part name="oid">1745863747437_JEffort</Part>
                <Part name="lineNumber">1</Part>
                <Part name="effortDate">Mon May 05 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1745863747438_JEffort">
                <Part name="effortExpense">75</Part>
                <Part name="effortAttr1">RIG-323</Part>
                <Part name="description">pair programming 
- fixing header logo size
- fixing footer logo visibility</Part>
                <Part name="oid">1745863747438_JEffort</Part>
                <Part name="lineNumber">2</Part>
                <Part name="effortDate">Mon May 05 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1745863747439_JEffort">
                <Part name="effortExpense">75</Part>
                <Part name="effortAttr1">RIG-437</Part>
                <Part name="description">- adding missing template changes on staging
- fixing lwc not setting fields</Part>
                <Part name="oid">1745863747439_JEffort</Part>
                <Part name="lineNumber">3</Part>
                <Part name="effortDate">Mon May 05 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1745863747440_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">RIG Daily</Part>
                <Part name="oid">1745863747440_JEffort</Part>
                <Part name="lineNumber">4</Part>
                <Part name="effortDate">Mon May 05 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1745863747441_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1"/>
                <Part name="description">coffee talk</Part>
                <Part name="oid">1745863747441_JEffort</Part>
                <Part name="lineNumber">5</Part>
                <Part name="effortDate">Mon May 05 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472173702_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">RIG Sync | Vorbereitung Testing Session</Part>
                <Part name="oid">1746472173702_JEffort</Part>
                <Part name="lineNumber">6</Part>
                <Part name="effortDate">Tue May 06 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472173703_JEffort">
                <Part name="effortExpense">90</Part>
                <Part name="effortAttr1">RIG-77</Part>
                <Part name="description">testing implementation</Part>
                <Part name="oid">1746472173703_JEffort</Part>
                <Part name="lineNumber">7</Part>
                <Part name="effortDate">Tue May 06 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472173704_JEffort">
                <Part name="effortExpense">180</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">Rigps Testing</Part>
                <Part name="oid">1746472173704_JEffort</Part>
                <Part name="lineNumber">8</Part>
                <Part name="effortDate">Tue May 06 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472173705_JEffort">
                <Part name="effortExpense">45</Part>
                <Part name="effortAttr1">RIG-412</Part>
                <Part name="description">creating new field, clarification naming, updated lwc with new input field</Part>
                <Part name="oid">1746472173705_JEffort</Part>
                <Part name="lineNumber">9</Part>
                <Part name="effortDate">Tue May 06 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472173706_JEffort">
                <Part name="effortExpense">150</Part>
                <Part name="effortAttr1">RIG-437</Part>
                <Part name="description">correcting field names, corected date format and vf apex element format</Part>
                <Part name="oid">1746472173706_JEffort</Part>
                <Part name="lineNumber">10</Part>
                <Part name="effortDate">Tue May 06 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472227027_JEffort">
                <Part name="effortExpense">375</Part>
                <Part name="effortAttr1">RIG-412</Part>
                <Part name="description">- created new fields
- updated lwc component in adressform site and checkout
-sync with Lauri</Part>
                <Part name="oid">1746472227027_JEffort</Part>
                <Part name="lineNumber">11</Part>
                <Part name="effortDate">Wed May 07 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472285053_JEffort">
                <Part name="effortExpense">75</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Informatik AG</Part>
                <Part name="oid">1746472285053_JEffort</Part>
                <Part name="lineNumber">12</Part>
                <Part name="effortDate">Wed May 07 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472285054_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1"/>
                <Part name="description">sync info AG</Part>
                <Part name="oid">1746472285054_JEffort</Part>
                <Part name="lineNumber">13</Part>
                <Part name="effortDate">Wed May 07 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472285055_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">RIG Daily</Part>
                <Part name="oid">1746472285055_JEffort</Part>
                <Part name="lineNumber">14</Part>
                <Part name="effortDate">Wed May 07 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472408556_JEffort">
                <Part name="effortExpense">375</Part>
                <Part name="effortAttr1">RIG-412</Part>
                <Part name="description">- updated shippingCalculator class
- updated CDG domain class
- updated CDG wrapper
- updated shipping address lwc</Part>
                <Part name="oid">1746472408556_JEffort</Part>
                <Part name="lineNumber">15</Part>
                <Part name="effortDate">Thu May 08 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472412550_JEffort">
                <Part name="effortExpense">45</Part>
                <Part name="effortAttr1">RIG-148</Part>
                <Part name="description">updated Registration Form</Part>
                <Part name="oid">1746472412550_JEffort</Part>
                <Part name="lineNumber">16</Part>
                <Part name="effortDate">Thu May 08 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472520484_JEffort">
                <Part name="effortExpense">45</Part>
                <Part name="effortAttr1"/>
                <Part name="description">update Skill sheet</Part>
                <Part name="oid">1746472520484_JEffort</Part>
                <Part name="lineNumber">17</Part>
                <Part name="effortDate">Fri May 09 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472520485_JEffort">
                <Part name="effortExpense">45</Part>
                <Part name="effortAttr1">RIG-412</Part>
                <Part name="description">updated domain class and lwc</Part>
                <Part name="oid">1746472520485_JEffort</Part>
                <Part name="lineNumber">18</Part>
                <Part name="effortDate">Fri May 09 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472520486_JEffort">
                <Part name="effortExpense">135</Part>
                <Part name="effortAttr1">RIG-445</Part>
                <Part name="description">updated templates with missing fields</Part>
                <Part name="oid">1746472520486_JEffort</Part>
                <Part name="lineNumber">19</Part>
                <Part name="effortDate">Fri May 09 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472520487_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">RIG-446</Part>
                <Part name="description">updated lwc to convert ids to uppercase</Part>
                <Part name="oid">1746472520487_JEffort</Part>
                <Part name="lineNumber">20</Part>
                <Part name="effortDate">Fri May 09 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472520488_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">RIG-447</Part>
                <Part name="description">added translations for delivery methods</Part>
                <Part name="oid">1746472520488_JEffort</Part>
                <Part name="lineNumber">21</Part>
                <Part name="effortDate">Fri May 09 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472520489_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">RIG Daily</Part>
                <Part name="oid">1746472520489_JEffort</Part>
                <Part name="lineNumber">22</Part>
                <Part name="effortDate">Fri May 09 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472520490_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1">D23-2</Part>
                <Part name="description">BreakoutSession</Part>
                <Part name="oid">1746472520490_JEffort</Part>
                <Part name="lineNumber">23</Part>
                <Part name="effortDate">Fri May 09 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472520724_JEffort">
                <Part name="effortExpense">60</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">finding solution for to many added to cart popup messages</Part>
                <Part name="oid">1746472520724_JEffort</Part>
                <Part name="lineNumber">24</Part>
                <Part name="effortDate">Fri May 09 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472630439_JEffort">
                <Part name="effortExpense">105</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Academy</Part>
                <Part name="oid">1746472630439_JEffort</Part>
                <Part name="lineNumber">25</Part>
                <Part name="effortDate">Mon May 12 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472630440_JEffort">
                <Part name="effortExpense">150</Part>
                <Part name="effortAttr1">RIG-148</Part>
                <Part name="description">added read permission and external credetial to guest profile
-field changes</Part>
                <Part name="oid">1746472630440_JEffort</Part>
                <Part name="lineNumber">26</Part>
                <Part name="effortDate">Mon May 12 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472647945_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1"/>
                <Part name="description">problem behandlung monitor und lan</Part>
                <Part name="oid">1746472647945_JEffort</Part>
                <Part name="lineNumber">27</Part>
                <Part name="effortDate">Mon May 12 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472647946_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">RIG Daily</Part>
                <Part name="oid">1746472647946_JEffort</Part>
                <Part name="lineNumber">28</Part>
                <Part name="effortDate">Mon May 12 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472693521_JEffort">
                <Part name="effortExpense">75</Part>
                <Part name="effortAttr1">RIG-408</Part>
                <Part name="description">checking possibilities to validate phone numbers
updating custom components</Part>
                <Part name="oid">1746472693521_JEffort</Part>
                <Part name="lineNumber">29</Part>
                <Part name="effortDate">Mon May 12 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472693522_JEffort">
                <Part name="effortExpense">75</Part>
                <Part name="effortAttr1">RIG-445</Part>
                <Part name="description">layout anpassungen</Part>
                <Part name="oid">1746472693522_JEffort</Part>
                <Part name="lineNumber">30</Part>
                <Part name="effortDate">Mon May 12 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472804726_JEffort">
                <Part name="effortExpense">45</Part>
                <Part name="effortAttr1">RIG-408</Part>
                <Part name="description">- adding basic validation</Part>
                <Part name="oid">1746472804726_JEffort</Part>
                <Part name="lineNumber">31</Part>
                <Part name="effortDate">Tue May 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472804727_JEffort">
                <Part name="effortExpense">120</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">Testrun</Part>
                <Part name="oid">1746472804727_JEffort</Part>
                <Part name="lineNumber">32</Part>
                <Part name="effortDate">Tue May 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472804728_JEffort">
                <Part name="effortExpense">120</Part>
                <Part name="effortAttr1">RIG-445</Part>
                <Part name="description">- fixing deployment issues
- changing layout of templates
- adding translations</Part>
                <Part name="oid">1746472804728_JEffort</Part>
                <Part name="lineNumber">33</Part>
                <Part name="effortDate">Tue May 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472804729_JEffort">
                <Part name="effortExpense">135</Part>
                <Part name="effortAttr1">RIG-148</Part>
                <Part name="description">-fixing deployment issues
- testing on staging
- update email template</Part>
                <Part name="oid">1746472804729_JEffort</Part>
                <Part name="lineNumber">34</Part>
                <Part name="effortDate">Tue May 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472804730_JEffort">
                <Part name="effortExpense">60</Part>
                <Part name="effortAttr1">RIG-412</Part>
                <Part name="description">testing on staging</Part>
                <Part name="oid">1746472804730_JEffort</Part>
                <Part name="lineNumber">35</Part>
                <Part name="effortDate">Tue May 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472804731_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">RIG Daily</Part>
                <Part name="oid">1746472804731_JEffort</Part>
                <Part name="lineNumber">36</Part>
                <Part name="effortDate">Tue May 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472804732_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Kickoff DSP Teamevent</Part>
                <Part name="oid">1746472804732_JEffort</Part>
                <Part name="lineNumber">37</Part>
                <Part name="effortDate">Tue May 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472900100_JEffort">
                <Part name="effortExpense">90</Part>
                <Part name="effortAttr1">RIG-152</Part>
                <Part name="description">- updated experience
- updated component</Part>
                <Part name="oid">1746472900100_JEffort</Part>
                <Part name="lineNumber">38</Part>
                <Part name="effortDate">Wed May 14 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472900101_JEffort">
                <Part name="effortExpense">45</Part>
                <Part name="effortAttr1">RIG-148</Part>
                <Part name="description">- updated experience</Part>
                <Part name="oid">1746472900101_JEffort</Part>
                <Part name="lineNumber">39</Part>
                <Part name="effortDate">Wed May 14 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472900102_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1">RIG-448</Part>
                <Part name="description">adjustment delivery/pickup date validation merge</Part>
                <Part name="oid">1746472900102_JEffort</Part>
                <Part name="lineNumber">40</Part>
                <Part name="effortDate">Wed May 14 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472900103_JEffort">
                <Part name="effortExpense">330</Part>
                <Part name="effortAttr1">RIG-445</Part>
                <Part name="description">- updated lwc and controller
- added custom field for billing name
- updated templates
- checked with lauri where re number gets set false</Part>
                <Part name="oid">1746472900103_JEffort</Part>
                <Part name="lineNumber">41</Part>
                <Part name="effortDate">Wed May 14 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746472900104_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">RIG Daily</Part>
                <Part name="oid">1746472900104_JEffort</Part>
                <Part name="lineNumber">42</Part>
                <Part name="effortDate">Wed May 14 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746473053734_JEffort">
                <Part name="effortExpense">90</Part>
                <Part name="effortAttr1">RIG-439</Part>
                <Part name="description">Refactoring validation check of form elements in checkout</Part>
                <Part name="oid">1746473053734_JEffort</Part>
                <Part name="lineNumber">43</Part>
                <Part name="effortDate">Fri May 16 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746473053735_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">RIG Daily</Part>
                <Part name="oid">1746473053735_JEffort</Part>
                <Part name="lineNumber">44</Part>
                <Part name="effortDate">Fri May 16 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746473053736_JEffort">
                <Part name="effortExpense">60</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Coffetalk</Part>
                <Part name="oid">1746473053736_JEffort</Part>
                <Part name="lineNumber">45</Part>
                <Part name="effortDate">Fri May 16 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1746473053737_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1">D23-2</Part>
                <Part name="description">BreakoutSession</Part>
                <Part name="oid">1746473053737_JEffort</Part>
                <Part name="lineNumber">46</Part>
                <Part name="effortDate">Fri May 16 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062210015_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Arbeitsplatz einrichten</Part>
                <Part name="oid">1749062210015_JEffort</Part>
                <Part name="lineNumber">47</Part>
                <Part name="effortDate">Tue Jun 10 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062210016_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Meeting eab, aufholen letzten 3 wochen, was ist passiert was ist der aktuelle stand, was sind offenen aufgaben</Part>
                <Part name="oid">1749062210016_JEffort</Part>
                <Part name="lineNumber">48</Part>
                <Part name="effortDate">Tue Jun 10 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062228474_JEffort">
                <Part name="effortExpense">105</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Requirements Engineering (Basics) Academy</Part>
                <Part name="oid">1749062228474_JEffort</Part>
                <Part name="lineNumber">49</Part>
                <Part name="effortDate">Tue Jun 10 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062235284_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1"/>
                <Part name="description">einlesen EKAG Miroboard</Part>
                <Part name="oid">1749062235284_JEffort</Part>
                <Part name="lineNumber">50</Part>
                <Part name="effortDate">Tue Jun 10 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062252863_JEffort">
                <Part name="effortExpense">75</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Ensuring the Billability of Services Academy</Part>
                <Part name="oid">1749062252863_JEffort</Part>
                <Part name="lineNumber">51</Part>
                <Part name="effortDate">Tue Jun 10 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062279920_JEffort">
                <Part name="effortExpense">120</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Software Quality (Advanced)</Part>
                <Part name="oid">1749062279920_JEffort</Part>
                <Part name="lineNumber">52</Part>
                <Part name="effortDate">Tue Jun 10 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062350628_JEffort">
                <Part name="effortExpense">135</Part>
                <Part name="effortAttr1">D23-28</Part>
                <Part name="description">anpassungs layout</Part>
                <Part name="oid">1749062350628_JEffort</Part>
                <Part name="lineNumber">53</Part>
                <Part name="effortDate">Fri May 16 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062351529_JEffort">
                <Part name="effortExpense">180</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Weiterbildung salesforce einstein</Part>
                <Part name="oid">1749062351529_JEffort</Part>
                <Part name="lineNumber">54</Part>
                <Part name="effortDate">Thu May 15 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062351824_JEffort">
                <Part name="effortExpense">195</Part>
                <Part name="effortAttr1">D23-28</Part>
                <Part name="description">refactoring von einigen unschönen stellen (HO)</Part>
                <Part name="oid">1749062351824_JEffort</Part>
                <Part name="lineNumber">55</Part>
                <Part name="effortDate">Tue Jun 10 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062390768_JEffort">
                <Part name="effortExpense">60</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Software Quality (Advanced)</Part>
                <Part name="oid">1749062390768_JEffort</Part>
                <Part name="lineNumber">56</Part>
                <Part name="effortDate">Wed Jun 11 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062390769_JEffort">
                <Part name="effortExpense">120</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Software Quality (Advanced)</Part>
                <Part name="oid">1749062390769_JEffort</Part>
                <Part name="lineNumber">57</Part>
                <Part name="effortDate">Wed Jun 11 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062390770_JEffort">
                <Part name="effortExpense">45</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Software Quality (Advanced)</Part>
                <Part name="oid">1749062390770_JEffort</Part>
                <Part name="lineNumber">58</Part>
                <Part name="effortDate">Wed Jun 11 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062390771_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">sync SPE</Part>
                <Part name="oid">1749062390771_JEffort</Part>
                <Part name="lineNumber">59</Part>
                <Part name="effortDate">Wed Jun 11 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062390772_JEffort">
                <Part name="effortExpense">120</Part>
                <Part name="effortAttr1">RIG-451</Part>
                <Part name="description">Figuring out a way on how to extend response count size for retrieved addresses in lwc</Part>
                <Part name="oid">1749062390772_JEffort</Part>
                <Part name="lineNumber">60</Part>
                <Part name="effortDate">Wed Jun 11 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062390773_JEffort">
                <Part name="effortExpense">60</Part>
                <Part name="effortAttr1">RIG-445</Part>
                <Part name="description">- recreating error cases
- checking updating process
- updated field updating</Part>
                <Part name="oid">1749062390773_JEffort</Part>
                <Part name="lineNumber">61</Part>
                <Part name="effortDate">Wed Jun 11 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062390774_JEffort">
                <Part name="effortExpense">60</Part>
                <Part name="effortAttr1">RIG-447</Part>
                <Part name="description">Figuring out a way to avoid rebuilding standard component for translating a single field</Part>
                <Part name="oid">1749062390774_JEffort</Part>
                <Part name="lineNumber">62</Part>
                <Part name="effortDate">Wed Jun 11 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062440833_JEffort">
                <Part name="effortExpense">135</Part>
                <Part name="effortAttr1">RIG-148</Part>
                <Part name="description">- added translations for phone number
- updated experience view
- added custom poperties for missing
translatable fields</Part>
                <Part name="oid">1749062440833_JEffort</Part>
                <Part name="lineNumber">63</Part>
                <Part name="effortDate">Thu May 15 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062440834_JEffort">
                <Part name="effortExpense">165</Part>
                <Part name="effortAttr1">RIG-408</Part>
                <Part name="description">- tried adding custom validation for phone number in
standard component in my Account view
- refactored way of setting Pickup option value as
Uppercase
- fixing merge errors</Part>
                <Part name="oid">1749062440834_JEffort</Part>
                <Part name="lineNumber">64</Part>
                <Part name="effortDate">Thu May 15 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062448999_JEffort">
                <Part name="effortExpense">90</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Software Quality (Advanced)</Part>
                <Part name="oid">1749062448999_JEffort</Part>
                <Part name="lineNumber">65</Part>
                <Part name="effortDate">Thu Jun 12 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062449000_JEffort">
                <Part name="effortExpense">105</Part>
                <Part name="effortAttr1">RIG-445</Part>
                <Part name="description">- changed type of Price fields to show missing zeros for all templates
- minor changes on styling in Summary table
- changed field label for Pickup
- added PackagingUnit to Amount value</Part>
                <Part name="oid">1749062449000_JEffort</Part>
                <Part name="lineNumber">66</Part>
                <Part name="effortDate">Thu Jun 12 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062471318_JEffort">
                <Part name="effortExpense">120</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Security Guideline</Part>
                <Part name="oid">1749062471318_JEffort</Part>
                <Part name="lineNumber">67</Part>
                <Part name="effortDate">Thu Jun 12 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062471319_JEffort">
                <Part name="effortExpense">45</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Software Quality (Advanced)</Part>
                <Part name="oid">1749062471319_JEffort</Part>
                <Part name="lineNumber">68</Part>
                <Part name="effortDate">Thu Jun 12 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062532796_JEffort">
                <Part name="effortExpense">90</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Security Guideline</Part>
                <Part name="oid">1749062532796_JEffort</Part>
                <Part name="lineNumber">69</Part>
                <Part name="effortDate">Fri Jun 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062532797_JEffort">
                <Part name="effortExpense">75</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Software Quality (Basics) Course</Part>
                <Part name="oid">1749062532797_JEffort</Part>
                <Part name="lineNumber">70</Part>
                <Part name="effortDate">Fri Jun 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062532798_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1">D23-2</Part>
                <Part name="description">BreakoutSession</Part>
                <Part name="oid">1749062532798_JEffort</Part>
                <Part name="lineNumber">71</Part>
                <Part name="effortDate">Fri Jun 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062534408_JEffort">
                <Part name="effortExpense">120</Part>
                <Part name="effortAttr1">EKP-90</Part>
                <Part name="description">EK Profishop | Sprint Planning</Part>
                <Part name="oid">1749062534408_JEffort</Part>
                <Part name="lineNumber">72</Part>
                <Part name="effortDate">Thu Jun 12 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062534601_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">EKP-90</Part>
                <Part name="description">EK Profishop | Daily Stand-up</Part>
                <Part name="oid">1749062534601_JEffort</Part>
                <Part name="lineNumber">73</Part>
                <Part name="effortDate">Fri Jun 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062571071_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">EKP-83</Part>
                <Part name="description">- Abklärung Inhalt für Doku zum Setup (NA?)</Part>
                <Part name="oid">1749062571071_JEffort</Part>
                <Part name="lineNumber">74</Part>
                <Part name="effortDate">Fri Jun 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062590757_JEffort">
                <Part name="effortExpense">90</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Security Guideline</Part>
                <Part name="oid">1749062590757_JEffort</Part>
                <Part name="lineNumber">75</Part>
                <Part name="effortDate">Fri Jun 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062590758_JEffort">
                <Part name="effortExpense">45</Part>
                <Part name="effortAttr1">EKP-90</Part>
                <Part name="description">- Einarbeitung/ Abklärung/ Überblick schaffen zu Projektstruktur, Packages, Confluence mit JMA</Part>
                <Part name="oid">1749062590758_JEffort</Part>
                <Part name="lineNumber">76</Part>
                <Part name="effortDate">Fri Jun 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062590759_JEffort">
                <Part name="effortExpense">120</Part>
                <Part name="effortAttr1">EKP-83</Part>
                <Part name="description">- Anlegen der Confluence Seite
- Erstellung grober Doku und Erfassung erster wichtiger Punkte</Part>
                <Part name="oid">1749062590759_JEffort</Part>
                <Part name="lineNumber">77</Part>
                <Part name="effortDate">Fri Jun 13 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062736905_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1">EKP-83</Part>
                <Part name="description">EK Profishop | Logging Standards</Part>
                <Part name="oid">1749062736905_JEffort</Part>
                <Part name="lineNumber">78</Part>
                <Part name="effortDate">Mon Jun 16 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062736906_JEffort">
                <Part name="effortExpense">45</Part>
                <Part name="effortAttr1">EKP-83</Part>
                <Part name="description">- updaten der Doku
- Szenarios und grundlegende Beispiele fälle und Standards eingefügt</Part>
                <Part name="oid">1749062736906_JEffort</Part>
                <Part name="lineNumber">79</Part>
                <Part name="effortDate">Mon Jun 16 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062736907_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">EKP-90</Part>
                <Part name="description">EK Profishop | Daily Stand-up</Part>
                <Part name="oid">1749062736907_JEffort</Part>
                <Part name="lineNumber">80</Part>
                <Part name="effortDate">Mon Jun 16 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062736908_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">RIG Daily</Part>
                <Part name="oid">1749062736908_JEffort</Part>
                <Part name="lineNumber">81</Part>
                <Part name="effortDate">Mon Jun 16 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062737289_JEffort">
                <Part name="effortExpense">315</Part>
                <Part name="effortAttr1">EKP-81</Part>
                <Part name="description">- Einrichtung Basis Setup PP mit MAA
- anlegen von Fonts, Farben und Button in Experience Builder</Part>
                <Part name="oid">1749062737289_JEffort</Part>
                <Part name="lineNumber">82</Part>
                <Part name="effortDate">Mon Jun 16 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062737420_JEffort">
                <Part name="effortExpense">60</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Security Guideline</Part>
                <Part name="oid">1749062737420_JEffort</Part>
                <Part name="lineNumber">83</Part>
                <Part name="effortDate">Mon Jun 16 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062791552_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1">EKP-83</Part>
                <Part name="description">- Komplettierung der Doku</Part>
                <Part name="oid">1749062791552_JEffort</Part>
                <Part name="lineNumber">84</Part>
                <Part name="effortDate">Tue Jun 17 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062791633_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">EKP-90</Part>
                <Part name="description">EK Profishop | Daily Stand-up</Part>
                <Part name="oid">1749062791633_JEffort</Part>
                <Part name="lineNumber">85</Part>
                <Part name="effortDate">Tue Jun 17 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062791634_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">Sync Tickets Rigips</Part>
                <Part name="oid">1749062791634_JEffort</Part>
                <Part name="lineNumber">86</Part>
                <Part name="effortDate">Tue Jun 17 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062791635_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1">D23-2</Part>
                <Part name="description">D23 | Statusupdate 1 Trimester 2/2025</Part>
                <Part name="oid">1749062791635_JEffort</Part>
                <Part name="lineNumber">87</Part>
                <Part name="effortDate">Tue Jun 17 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062813765_JEffort">
                <Part name="effortExpense">330</Part>
                <Part name="effortAttr1">EKP-81</Part>
                <Part name="description">- Experience Daten im Projekt Packages verschieben
- Branding anpassen
- Klärung von Styleguide Fragen wegen doppelter Angaben für Elemente
- anlegen restlicher Variablen im General Stylesheet
- Dokumentation überprüfen</Part>
                <Part name="oid">1749062813765_JEffort</Part>
                <Part name="lineNumber">88</Part>
                <Part name="effortDate">Tue Jun 17 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749062860808_JEffort">
                <Part name="effortExpense">60</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Security Guideline</Part>
                <Part name="oid">1749062860808_JEffort</Part>
                <Part name="lineNumber">89</Part>
                <Part name="effortDate">Tue Jun 17 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063102421_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">EKP-90</Part>
                <Part name="description">EK Profishop | Daily Stand-up</Part>
                <Part name="oid">1749063102421_JEffort</Part>
                <Part name="lineNumber">90</Part>
                <Part name="effortDate">Thu Jun 19 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063102422_JEffort">
                <Part name="effortExpense">150</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Pratice Summit</Part>
                <Part name="oid">1749063102422_JEffort</Part>
                <Part name="lineNumber">91</Part>
                <Part name="effortDate">Thu Jun 19 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063102598_JEffort">
                <Part name="effortExpense">360</Part>
                <Part name="effortAttr1">EKP-81</Part>
                <Part name="description">- Erweiterung Component Library
   -  Checkboxen
   - Radio Buttons
   - Language Switcher
   - Input Fields</Part>
                <Part name="oid">1749063102598_JEffort</Part>
                <Part name="lineNumber">92</Part>
                <Part name="effortDate">Thu Jun 19 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063102994_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">EKP-90</Part>
                <Part name="description">EK Profishop | Daily Stand-up</Part>
                <Part name="oid">1749063102994_JEffort</Part>
                <Part name="lineNumber">93</Part>
                <Part name="effortDate">Wed Jun 18 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063102995_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1">RIG-113</Part>
                <Part name="description">Sync Tickts</Part>
                <Part name="oid">1749063102995_JEffort</Part>
                <Part name="lineNumber">94</Part>
                <Part name="effortDate">Wed Jun 18 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063102996_JEffort">
                <Part name="effortExpense">60</Part>
                <Part name="effortAttr1">D23-2</Part>
                <Part name="description">Checkpoint Gespräch</Part>
                <Part name="oid">1749063102996_JEffort</Part>
                <Part name="lineNumber">95</Part>
                <Part name="effortDate">Wed Jun 18 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063102997_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Coffee Talk VWA JUC</Part>
                <Part name="oid">1749063102997_JEffort</Part>
                <Part name="lineNumber">96</Part>
                <Part name="effortDate">Wed Jun 18 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063103248_JEffort">
                <Part name="effortExpense">510</Part>
                <Part name="effortAttr1">EKP-81</Part>
                <Part name="description">- Component Library erstellen, Stylen
-  Styling Setup
- Pipeline fix
- Fix Merge Conflicts</Part>
                <Part name="oid">1749063103248_JEffort</Part>
                <Part name="lineNumber">97</Part>
                <Part name="effortDate">Wed Jun 18 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063104896_JEffort">
                <Part name="effortExpense">480</Part>
                <Part name="effortAttr1"/>
                <Part name="description">dotSource Firmenoffsite 2025</Part>
                <Part name="oid">1749063104896_JEffort</Part>
                <Part name="lineNumber">98</Part>
                <Part name="effortDate">Fri Jun 20 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063242717_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1">EKP-83</Part>
                <Part name="description">- beispiel bilder hinzugefügt</Part>
                <Part name="oid">1749063242717_JEffort</Part>
                <Part name="lineNumber">99</Part>
                <Part name="effortDate">Mon Jun 23 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063242718_JEffort">
                <Part name="effortExpense">90</Part>
                <Part name="effortAttr1">EKP-81</Part>
                <Part name="description">- anpassung inputs/dropdowns</Part>
                <Part name="oid">1749063242718_JEffort</Part>
                <Part name="lineNumber">100</Part>
                <Part name="effortDate">Mon Jun 23 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063242719_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">EKP-90</Part>
                <Part name="description">EK Profishop | Daily Stand-up</Part>
                <Part name="oid">1749063242719_JEffort</Part>
                <Part name="lineNumber">101</Part>
                <Part name="effortDate">Mon Jun 23 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063266108_JEffort">
                <Part name="effortExpense">240</Part>
                <Part name="effortAttr1">EKP-9</Part>
                <Part name="description">- einfügen von custom HTML zu anzeigen weiterer Felder
- anpassung der neuen Feldnamen</Part>
                <Part name="oid">1749063266108_JEffort</Part>
                <Part name="lineNumber">102</Part>
                <Part name="effortDate">Mon Jun 23 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063278945_JEffort">
                <Part name="effortExpense">150</Part>
                <Part name="effortAttr1">EKP-41</Part>
                <Part name="description">- Anpassung Style der Passwort vergessen Seite</Part>
                <Part name="oid">1749063278945_JEffort</Part>
                <Part name="lineNumber">103</Part>
                <Part name="effortDate">Mon Jun 23 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063280809_JEffort">
                <Part name="effortExpense">45</Part>
                <Part name="effortAttr1"/>
                <Part name="description">Nachtrag BLOG</Part>
                <Part name="oid">1749063280809_JEffort</Part>
                <Part name="lineNumber">104</Part>
                <Part name="effortDate">Mon Jun 23 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063347768_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">EKP-90</Part>
                <Part name="description">Refinement</Part>
                <Part name="oid">1749063347768_JEffort</Part>
                <Part name="lineNumber">105</Part>
                <Part name="effortDate">Tue Jun 24 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063347769_JEffort">
                <Part name="effortExpense">435</Part>
                <Part name="effortAttr1">EKP-41</Part>
                <Part name="description">- Anpassung-Passwort-vergessen-Seite
- Anpassung Anzeige Kundendaten
 - Anpassung-Shop-Logos
- Anpassung Login Seite
- Anpassung Font-Size viewports</Part>
                <Part name="oid">1749063347769_JEffort</Part>
                <Part name="lineNumber">106</Part>
                <Part name="effortDate">Tue Jun 24 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063347770_JEffort">
                <Part name="effortExpense">15</Part>
                <Part name="effortAttr1">EKP-90</Part>
                <Part name="description">EK Profishop | Daily Stand-up</Part>
                <Part name="oid">1749063347770_JEffort</Part>
                <Part name="lineNumber">107</Part>
                <Part name="effortDate">Tue Jun 24 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1749063347771_JEffort">
                <Part name="effortExpense">30</Part>
                <Part name="effortAttr1">D23-2</Part>
                <Part name="description">Coffe Talk</Part>
                <Part name="oid">1749063347771_JEffort</Part>
                <Part name="lineNumber">108</Part>
                <Part name="effortDate">Tue Jun 24 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getEffortsByOidAndDateRangeLoader</Part>
            </Record>
        </ds:getEffortsByOidAndDateRangeResponse>
    </S:Body>
</S:Envelope>`

var Appointment_response = `<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://www.w3.org/2003/05/soap-envelope">
    <S:Body>
        <ds:getAppointmentsByOidAndDateRangeResponse xmlns:ds="http://www.dotsource.de">
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1707851602760_JAppointment">
                <Part name="eventEndDate">Fri May 23 00:00:00 CEST 2025</Part>
                <Part name="eventType">noBudget</Part>
                <Part name="oid">1707851602760_JAppointment</Part>
                <Part name="lineNumber">1</Part>
                <Part name="eventStartDate">Mon May 19 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getAppointmentsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1707851602919_JAppointment">
                <Part name="eventEndDate">Fri Jun 06 00:00:00 CEST 2025</Part>
                <Part name="eventType">noBudget</Part>
                <Part name="oid">1707851602919_JAppointment</Part>
                <Part name="lineNumber">2</Part>
                <Part name="eventStartDate">Mon Jun 02 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getAppointmentsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1739475565298_JAppointment">
                <Part name="eventEndDate">Fri May 02 00:00:00 CEST 2025</Part>
                <Part name="eventType">baseOrRemainingBudget</Part>
                <Part name="oid">1739475565298_JAppointment</Part>
                <Part name="lineNumber">3</Part>
                <Part name="eventStartDate">Fri May 02 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getAppointmentsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1739475565888_JAppointment">
                <Part name="eventEndDate">Fri May 30 00:00:00 CEST 2025</Part>
                <Part name="eventType">baseOrRemainingBudget</Part>
                <Part name="oid">1739475565888_JAppointment</Part>
                <Part name="lineNumber">4</Part>
                <Part name="eventStartDate">Fri May 30 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getAppointmentsByOidAndDateRangeLoader</Part>
            </Record>
            <Record xmlns="http://www.projektron.de/bcs/processing/data" oid="1747736619894_JAppointment">
                <Part name="eventEndDate">Wed May 28 00:00:00 CEST 2025</Part>
                <Part name="eventType">sickness</Part>
                <Part name="oid">1747736619894_JAppointment</Part>
                <Part name="lineNumber">5</Part>
                <Part name="eventStartDate">Mon May 26 00:00:00 CEST 2025</Part>
                <Part name="%loadername%">getAppointmentsByOidAndDateRangeLoader</Part>
            </Record>
        </ds:getAppointmentsByOidAndDateRangeResponse>
    </S:Body>
</S:Envelope>`
