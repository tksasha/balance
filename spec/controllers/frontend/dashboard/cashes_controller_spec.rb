# frozen_string_literal: true

RSpec.describe Frontend::Dashboard::CashesController do
  it { is_expected.to be_a(BaseController) }

  describe '#resource' do
    context 'when @resource is set' do
      before { subject.instance_variable_set :@resource, :resource }

      its(:resource) { is_expected.to eq :resource }
    end

    context 'when @resource is not set' do
      let(:params) { { id: 28 } }

      before do
        allow(subject).to receive(:params).and_return(params)

        allow(Cash).to receive(:find).with(28).and_return(:resource)
      end

      its(:resource) { is_expected.to eq :resource }
    end
  end

  describe '#resource_params' do
    let(:params) { acp(cash: { name: nil, formula: nil }) }

    before { allow(subject).to receive(:params).and_return(params) }

    its(:resource_params) { is_expected.to eq params.require(:cash).permit! }
  end
end
