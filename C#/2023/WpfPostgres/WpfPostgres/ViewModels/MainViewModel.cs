using GalaSoft.MvvmLight;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace WpfPostgres.ViewModels
{
    public class MainViewModel : ViewModelBase
    {
        public ProductViewModel Product { get; set; }
        public ButtonsViewModel Buttons { get; set; }

    }
}
