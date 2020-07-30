<h1>SOCAT ALTER</h1>
<p>
This is a small application for fowarding port from your localhost to public IP.
<br/>
Just specific your source->destination port in <code>routes.json</code> file.
<br/>
Example:
<pre>
[
    {
        "src":"2375",
        "dst":"2375"
    },
    {
        "src":"8080",
        "dst":"80"
    }
]
</pre>
</p>
<p>
<b>NOTE: </b><i>This is a private application, use it in a safe local network only. I'm not take responsible for any data leak from this application.</i>
<p>
