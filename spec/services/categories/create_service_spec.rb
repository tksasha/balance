# frozen_string_literal: true

RSpec.describe Categories::CreateService do
  subject { described_class.new params }

  let(:params) { acp(category: { name: nil, income: nil, visible: nil, currency: nil }) }

  its(:resource_params) { is_expected.to eq params.require(:category).permit! }

  describe '#category' do
    context do
      before { subject.instance_variable_set :@category, :category }

      its(:category) { is_expected.to eq :category }
    end

    context do
      before do
        allow(subject).to receive(:resource_params).and_return(:resource_params)

        allow(Category).to receive(:new).with(:resource_params).and_return(:category)
      end

      its(:category) { is_expected.to eq :category }
    end
  end

  describe '#call' do
    let(:category) { stub_model Category }

    before { allow(subject).to receive(:category).and_return(category) }

    context do
      before { allow(category).to receive(:save).and_return(true) }

      its(:call) { is_expected.to be_success }

      its('call.object') { is_expected.to eq category }
    end

    context do
      before { allow(category).to receive(:save).and_return(false) }

      its(:call) { is_expected.to be_failure }

      its('call.object') { is_expected.to eq category }
    end
  end
end
