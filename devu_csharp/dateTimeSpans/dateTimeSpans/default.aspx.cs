using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace dateTimeSpans
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

        }

        protected void okButton_Click(object sender, EventArgs e)
        {
            // TimeSpan myTimespan = TimeSpan.Parse("10.1:05:10.01");

            // resultLabel.Text = myTimespan.ToString();

            //http://is.gd/timespan

            //TimeSpan myValue = new TimeSpan(10,1,05,10,01);

            //resultLabel.Text = myValue.ToString();

            DateTime meuAniversario = DateTime.Parse("06/12/1973");
            TimeSpan minhaIdade = DateTime.Now.Subtract(meuAniversario);
            
            resultLabel.Text = minhaIdade.ToString(); 
        }
    }
}