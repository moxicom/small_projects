using GalaSoft.MvvmLight;
using GalaSoft.MvvmLight.Command;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Input;

namespace WpfPostgres.ViewModels
{
    public class MainViewModelPages: ViewModelBase, INotifyPropertyChanged
    {
        private ViewModelBase _currentViewModel;

        public ViewModelBase CurrentViewModel
        {
            get { return _currentViewModel; }
            set
            {
                if (_currentViewModel != value)
                {
                    _currentViewModel = value;
                    OnPropertyChanged(nameof(CurrentViewModel));
                }
            }
        }

        public ICommand HomeBtnClick { get; }
        public ICommand AccountBtnClick { get; }

        public event PropertyChangedEventHandler? PropertyChanged;

        public MainViewModelPages()
        {
            CurrentViewModel = new HomeViewModel();
            HomeBtnClick = new RelayCommand(OpenHome);
            AccountBtnClick = new RelayCommand(OpenAccount);
        }

        public void OnPropertyChanged([CallerMemberName] string prop = "")
        {
            if (PropertyChanged != null)
                PropertyChanged(this, new PropertyChangedEventArgs(prop));
        }

        private void OpenHome()
        {
            CurrentViewModel = new HomeViewModel();
        }

        private void OpenAccount()
        {
            CurrentViewModel = new AccountViewModel();
        }


    }
}
