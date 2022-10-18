# frozen_string_literal: true

RSpec.describe Frontend::Dashboard::CashesController, type: :controller do
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

  describe '#dashboard' do
    before do
      allow(subject).to receive(:params).and_return(:params)

      allow(Frontend::Dashboard).to receive(:new).with(:params).and_return(:dashboard)
    end

    its(:dashboard) { is_expected.to eq :dashboard }

    its(:_helper_methods) { is_expected.to include :dashboard }
  end
end
