using GalaSoft.MvvmLight.Command;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Security.RightsManagement;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Input;

namespace WpfPostgres.ViewModels
{
    public class ButtonsViewModel : INotifyPropertyChanged
    {
        private string _Btn1Content = "Click me!";
        private string _Btn2Content = "Click me!";

        public ICommand FirstButtonCommand { get; private set; }
        public ICommand SecondButtonCommand { get; private set; }


        public ButtonsViewModel()
        {
            FirstButtonCommand = new RelayCommand(ExecuteFirstButton);
            SecondButtonCommand = new RelayCommand(ExecuteSecondButton);
        }

        public event PropertyChangedEventHandler? PropertyChanged;

        public void OnPropertyChanged([CallerMemberName] string prop = "")
        {
            if (PropertyChanged != null)
                PropertyChanged(this, new PropertyChangedEventArgs(prop));
        }

        public string Btn1Content
        {
            get { return _Btn1Content; }
            set
            {
                if (_Btn1Content != value)
                {
                    _Btn1Content = value;
                    OnPropertyChanged(nameof(Btn1Content));
                }
            }
        }

        public string Btn2Content
        {
            get { return _Btn2Content; }
            set
            {
                if (_Btn2Content != value)
                {
                    _Btn2Content = value;
                    OnPropertyChanged(nameof(Btn2Content));
                }
            }
        }

        public void ExecuteFirstButton()
        {
            Btn1Content = "11";
        }

        public void ExecuteSecondButton()
        {
            Btn2Content = "22";
        }
    }
}
