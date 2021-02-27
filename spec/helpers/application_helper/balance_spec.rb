# frozen_string_literal: true

RSpec.describe ApplicationHelper, type: :helper do
  subject { helper }

  describe '#balance' do
    context do
      before { subject.instance_variable_set :@balance, :balance }

      its(:balance) { should eq :balance }
    end

    context do
      let(:params) { double }

      before { allow(subject).to receive(:params).and_return(params) }

      before { allow(CalculateBalanceService).to receive(:calculate).with(params).and_return(22.15) }

      its(:balance) { should eq 22.15 }
    end
  end
end
