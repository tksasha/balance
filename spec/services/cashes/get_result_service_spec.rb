# frozen_string_literal: true

RSpec.describe Cashes::GetResultService, type: :service do
  subject { described_class.new action_name, params }

  let(:params) { double }

  describe '#call' do
    context do
      let(:action_name) { 'show' }

      before { allow(Cashes::GetResourceService).to receive(:call).with(params).and_return(:result) }

      its(:call) { should eq :result }
    end

    context do
      let(:action_name) { 'edit' }

      before { allow(Cashes::GetResourceService).to receive(:call).with(params).and_return(:result) }

      its(:call) { should eq :result }
    end

    context do
      let(:action_name) { 'update' }

      before { allow(Cashes::UpdateService).to receive(:call).with(params).and_return(:result) }

      its(:call) { should eq :result }
    end
  end
end