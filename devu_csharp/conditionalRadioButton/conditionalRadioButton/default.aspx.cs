using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace conditionalRadioButton
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

        }

        protected void okButton_Click(object sender, EventArgs e)
        {
            if (pencilRadioButton.Checked)
            {
                resultLabel.Text = "Pencil.";
                resultimage.ImageUrl = "pencil.jpg";

            }
            else if (penRadioButton.Checked)
            {

                resultLabel.Text = "Pen.";
                resultimage.ImageUrl = "pen.jpg";
            }
            else if (phoneRadioButton.Checked)
            {
                resultLabel.Text = "Phone.";
                resultimage.ImageUrl = "phone.jpg";
            }
            else if (tabletRadioButton.Checked)
            {
                resultLabel.Text = "Tablet.";
                resultimage.ImageUrl = "tablet.jpg";
            }
            else 
            {
                resultLabel.Text = "Select an option.";
                
            }
        }
    }
}