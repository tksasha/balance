# frozen_string_literal: true

RSpec.describe Items::GetResultService do
  subject { described_class.new action_name, params }

  let(:params) { double }

  describe '#call' do
    context do
      let(:action_name) { 'create' }

      before { allow(Items::CreateService).to receive(:call).with(params).and_return(:result) }

      its(:call) { should eq :result }
    end

    context do
      let(:action_name) { 'edit' }

      before { allow(Items::GetResourceService).to receive(:call).with(params).and_return(:result) }

      its(:call) { should eq :result }
    end
  end
end
