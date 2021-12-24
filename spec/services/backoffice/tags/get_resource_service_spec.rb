# frozen_string_literal: true

RSpec.describe Backoffice::Tags::GetResourceService do
  subject { described_class.new params }

  let(:params) { { id: 39 } }

  describe '#tag' do
    context do
      before { subject.instance_variable_set :@tag, :tag }

      its(:tag) { should eq :tag }
    end

    context do
      before { allow(Tag).to receive(:find).with(39).and_return(:tag) }

      its(:tag) { should eq :tag }
    end
  end

  describe '#call' do
    before { allow(subject).to receive(:tag).and_return(:tag) }

    its(:call) { should be_success }

    its('call.object') { should eq :tag }
  end
end
