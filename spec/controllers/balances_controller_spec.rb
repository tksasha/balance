# frozen_string_literal: true

RSpec.describe BalancesController, type: :controller do
  describe '#resource' do
    context do
      before { subject.instance_variable_set :@resource, :resource }

      its(:resource) { should eq :resource }
    end

    context do
      let(:params) { double }

      before { expect(subject).to receive(:params).and_return(params) }

      before { expect(BalanceService).to receive(:new).with(params).and_return(:resource) }

      its(:resource) { should eq :resource }
    end
  end
end
